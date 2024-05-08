package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"bridge/lib"
)

type Message struct {
	EncodedSwap string
	Sign        string
	Recipient   string
	Initiator   string
	TargetChain string
}

type HttpResponse struct {
	// 0. SUCCESS 1.ERROR 2.PENDING
	Status int    `json:"status"`
	Err    string `json:"err"`
}

const (
	ONE_MESON string = "0x4f0F26939BB50611124274E2051c42825130a1E8"
	ETH_MESON string = "0x483c1FE0E455912A69De699DC967b6d0E1e4f94a"

	ONE_RPC string = "https://rpc.devnet.onenesslabs.io"
	ETH_RPC string = "https://sepolia.drpc.org"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	response := HttpResponse{
		Status: 1,
		Err:    "",
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("not found param, body:", r.Body)

		response.Err = "not found param"
		jsonStr, _ := json.Marshal(response)
		fmt.Fprintf(w, string(jsonStr))
		return
	}
	var m Message
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println("unmarshal body fail:", body)

		response.Err = "unmarshal body fail"
		jsonStr, _ := json.Marshal(response)
		fmt.Fprintf(w, string(jsonStr))
		return
	}
	log.Printf("param:%+v\n", m)
	err = callBridge(m)
	if err != nil {
		log.Println("callBridge error:", err.Error())

		response.Err = err.Error()
		jsonStr, _ := json.Marshal(response)
		fmt.Fprintf(w, string(jsonStr))
		return

	} else {
		response.Status = 0
		jsonStr, _ := json.Marshal(response)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

}

func getAuth(client *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	targetNonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	targetGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(targetNonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = targetGasPrice
	return auth, nil
}

func callBridge(m Message) error {
	var targetMesonAddr string

	var targetRpcUrl string
	var originMesonAddr string
	var originRpcUrl string
	if m.TargetChain == "eth" {
		targetMesonAddr = ETH_MESON
		targetRpcUrl = ETH_RPC
		originMesonAddr = ONE_MESON
		originRpcUrl = ONE_RPC
	} else {
		targetMesonAddr = ONE_MESON
		targetRpcUrl = ONE_RPC

		originMesonAddr = ETH_MESON
		originRpcUrl = ETH_RPC
	}

	_ = originMesonAddr
	_ = originRpcUrl

	targetMeson := common.HexToAddress(targetMesonAddr)

	targetClient, err := ethclient.Dial(targetRpcUrl)
	if err != nil {
		log.Println("connect target rpc:", err, targetRpcUrl)
		return err
	}

	targetInstance, err := lib.NewMeson(targetMeson, targetClient)
	if err != nil {
		log.Println("new Meson:", err)
		return err
	}

	_ = targetClient

	privateKey, err := crypto.HexToECDSA(goDotEnvVariable("PRIVATE_KEY"))
	if err != nil {
		log.Println("get prikey:", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return errors.New("annot assert type: publicKey is not of type *ecdsa.PublicKey")

	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("bridgeProxy wallet address:%s\n", fromAddress)

	encodedSwap := new(big.Int)
	fmt.Println(strings.TrimPrefix(m.EncodedSwap, "0x"))
	encodedSwap.SetString(strings.TrimPrefix(m.EncodedSwap, "0x"), 16)
	hexSign := strings.TrimPrefix(m.Sign, "0x")

	byteArray, err := hex.DecodeString(hexSign)

	var r [32]byte
	var vs [32]byte
	copy(r[:], byteArray[:32])
	copy(vs[:], byteArray[32:64])

	recipient := common.HexToAddress(m.Recipient)
	initiator := common.HexToAddress(m.Initiator)

	// send lockswap to target chain
	auth, err := getAuth(targetClient, fromAddress, privateKey)
	if err != nil {
		log.Println("NewKeyedTransactor:", err)
		return err
	}

	lockTx, err := targetInstance.LockSwap(auth, encodedSwap, initiator)
	if err != nil {
		log.Println("LockSwap error:", err)
		return err
	}

	log.Printf("LockSwap %s", lockTx.Hash().Hex())

	//send release tx to target chain

	auth, err = getAuth(targetClient, fromAddress, privateKey)
	if err != nil {
		log.Println("NewKeyedTransactor:", err)
		return err
	}

	releaseTx, err := targetInstance.Release(auth, encodedSwap, r, vs, initiator, recipient)
	if err != nil {
		log.Printf("releaseTx error: %s\n", err)
		return err
	}

	log.Println("releaseTx sent: %s", releaseTx.Hash().Hex())

	//send executeSwap tx to origin chain

	/*originMeson := common.HexToAddress(originMesonAddr)
	originClient, err := ethclient.Dial(originRpcUrl)
	if err != nil {
		log.Println(err)
		return err
	}

	originInstance, err := lib.NewMeson(originMeson, originClient)
	// log.Printf("originInstance %s error:%s\n", originInstance, err)
	_ = originInstance

	auth, err = getAuth(originClient, fromAddress, privateKey)
	if err != nil {
		log.Println(err)
		return err
	}

	executeSwapTx, err := originInstance.DirectExecuteSwap(auth, encodedSwap, r, vs, initiator, recipient)
	if err != nil {
		log.Printf("executeSwap error: %s\n", err)
		return err
	}

	fmt.Printf("executeSwap sent: %s", executeSwapTx.Hash().Hex())*/

	return nil
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// func main() {
// 	// callBridge()
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", Handler)

// 	err := http.ListenAndServe(":3333", mux)
// 	if errors.Is(err, http.ErrServerClosed) {
// 		fmt.Printf("server closed\n")
// 	} else if err != nil {
// 		fmt.Printf("error starting server: %s\n", err)
// 		os.Exit(1)
// 	}

// }
