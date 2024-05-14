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
	"time"

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

const AttemptLimit = 3
const (
	ONE_MESON   string = "0x4f0F26939BB50611124274E2051c42825130a1E8"
	ETH_MESON   string = "0x483c1FE0E455912A69De699DC967b6d0E1e4f94a"
	ETH_CHAINID int    = 11155111

	ONE_RPC     string = "https://rpc.devnet.onenesslabs.io"
	ETH_RPC     string = "https://ethereum-sepolia-rpc.publicnode.com"
	ONE_CHAINID int    = 123666
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

func getAuth(client *ethclient.Client, fromAddress common.Address, privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
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
	log.Println("chainId:", chainId)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Println("NewKeyedTransactorWithChainID error:", err)
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(targetNonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = targetGasPrice
	return auth, nil
}

func getReceiptByTxhash(txhash string, rpcclient *ethclient.Client) bool {

	_txHash := common.HexToHash(txhash)
	receipt, err := rpcclient.TransactionReceipt(context.Background(), _txHash)

	if err != nil {
		log.Printf("TransactionReceipt: %s\n", err)
		return false
	}
	log.Printf("txhash:%s receipt blocknum:%s\n", txhash, receipt.BlockNumber)
	return true

}

func checkTx(txhash string, rpcClient *ethclient.Client, sleep int) bool {

	for i := 1; i <= AttemptLimit; i++ {
		time.Sleep(time.Second * time.Duration(sleep))
		ok := getReceiptByTxhash(txhash, rpcClient)
		if ok {
			return true
		}
	}
	return false
}

func callBridge(m Message) error {
	var targetMesonAddr string

	var targetRpcUrl string
	var originMesonAddr string
	var originRpcUrl string
	var originChainId *big.Int
	var targetChainId *big.Int
	var targetSleep int
	var originSleep int
	if m.TargetChain == "eth" {
		targetMesonAddr = ETH_MESON
		targetRpcUrl = ETH_RPC
		targetChainId = new(big.Int).SetUint64(uint64(ETH_CHAINID))
		_ = targetChainId

		originMesonAddr = ONE_MESON
		originRpcUrl = ONE_RPC
		originChainId = new(big.Int).SetUint64(uint64(ONE_CHAINID))

		targetSleep = 15
		originSleep = 3

	} else {
		targetMesonAddr = ONE_MESON
		targetRpcUrl = ONE_RPC
		targetChainId = new(big.Int).SetUint64(uint64(ONE_CHAINID))
		_ = targetChainId

		originMesonAddr = ETH_MESON
		originRpcUrl = ETH_RPC
		originChainId = new(big.Int).SetUint64(uint64(ETH_CHAINID))

		targetSleep = 3
		originSleep = 15
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
	_ = targetInstance
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
	_ = recipient
	initiator := common.HexToAddress(m.Initiator)
	_ = initiator
	// send lockswap to origin chain
	auth, err := getAuth(targetClient, fromAddress, privateKey, targetChainId)
	if err != nil {
		log.Println("NewKeyedTransactor:", err)
		return err
	}

	lockTx, err := targetInstance.LockSwap(auth, encodedSwap, initiator)
	if err != nil {
		log.Println("LockSwap error:", err)
		return err
	}

	log.Printf("LockSwap %s\n", lockTx.Hash().Hex())

	if !checkTx(lockTx.Hash().Hex(), targetClient, targetSleep) {
		return errors.New("lockSwap tx failed")
	}

	//send release tx to target chain

	auth, err = getAuth(targetClient, fromAddress, privateKey, targetChainId)
	if err != nil {
		log.Println("NewKeyedTransactor:", err)
		return err
	}

	releaseTx, err := targetInstance.Release(auth, encodedSwap, r, vs, initiator, recipient)
	if err != nil {
		log.Printf("releaseTx error: %s\n", err)
		return err
	}

	log.Printf("releaseTx sent: %s\n", releaseTx.Hash().Hex())

	if !checkTx(releaseTx.Hash().Hex(), targetClient, targetSleep) {
		return errors.New("release tx failed")
	}

	//send executeSwap tx to origin chain

	originMeson := common.HexToAddress(originMesonAddr)
	originClient, err := ethclient.Dial(originRpcUrl)
	if err != nil {
		log.Println(err)
		return err
	}

	originInstance, err := lib.NewMeson(originMeson, originClient)
	// log.Printf("originInstance %s error:%s\n", originInstance, err)
	_ = originInstance

	auth, err = getAuth(originClient, fromAddress, privateKey, originChainId)
	if err != nil {
		log.Println(err)
		return err
	}

	executeSwapTx, err := originInstance.ExecuteSwap(auth, encodedSwap, r, vs, initiator, false)
	if err != nil {
		log.Printf("executeSwapTx error: %s\n", err)
		return err
	}

	log.Printf("executeSwapTx sent: %s\n", executeSwapTx.Hash().Hex())

	if !checkTx(executeSwapTx.Hash().Hex(), originClient, originSleep) {
		return errors.New("executeSwap tx failed")
	}

	return nil
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
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
