package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Message struct {
	EncodedSwap string
	Sign        string
	Recipient   string
	Initiator   string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// fmt.Printf("could not read body: %s\n", err)
		fmt.Fprintf(w, "not found param")
		return
	}
	var m Message
	err = json.Unmarshal(body, &m)
	if err != nil {

		fmt.Fprintf(w, "unmarshal body fail")
		return
	}
	fmt.Println(m)
	// err = callBridge(m)
	err = nil
	if err != nil {
		fmt.Fprintf(w, "fail")
	} else {
		fmt.Fprintf(w, "success")
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

// func callBridge(m Message) error {
// 	targetMesonAddr := common.HexToAddress("0x4f0F26939BB50611124274E2051c42825130a1E8")

// 	targetClient, err := ethclient.Dial("https://rpc.devnet.onenesslabs.io")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	originMesonAddr := common.HexToAddress("0x483c1FE0E455912A69De699DC967b6d0E1e4f94a")
// 	originClient, err := ethclient.Dial("https://sepolia.drpc.org")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	originInstance, err := lib.NewMeson(originMesonAddr, originClient)
// 	// log.Printf("originInstance %s error:%s\n", originInstance, err)
// 	_ = originInstance

// 	targetInstance, err := lib.NewMeson(targetMesonAddr, targetClient)
// 	// log.Printf("targetInstance %s error:%s\n", targetInstance, err)

// 	_ = targetClient

// 	privateKey, err := crypto.HexToECDSA(goDotEnvVariable("PRIVATE_KEY"))
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 		return errors.New("annot assert type: publicKey is not of type *ecdsa.PublicKey")

// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	fmt.Printf("bridgeProxy wallet address:%s\n", fromAddress)

// 	encodedSwap := new(big.Int)
// 	fmt.Println(strings.TrimPrefix(m.EncodedSwap, "0x"))
// 	encodedSwap.SetString(strings.TrimPrefix(m.EncodedSwap, "0x"), 16)
// 	hexSign := strings.TrimPrefix(m.Sign, "0x")

// 	byteArray, err := hex.DecodeString(hexSign)

// 	var r [32]byte
// 	var vs [32]byte
// 	copy(r[:], byteArray[:32])
// 	copy(vs[:], byteArray[32:64])

// 	recipient := common.HexToAddress(m.Recipient)
// 	initiator := common.HexToAddress(m.Initiator)

// 	// send lockswap to target chain
// 	auth, err := getAuth(targetClient, fromAddress, privateKey)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	lockTx, err := targetInstance.LockSwap(auth, encodedSwap, initiator)
// 	if err != nil {
// 		log.Printf("lockTx error: %s\n", err)
// 		return err
// 	}

// 	log.Printf("lockTx %s", lockTx.Hash().Hex())

// 	//send release tx to target chain

// 	auth, err = getAuth(targetClient, fromAddress, privateKey)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	releaseTx, err := targetInstance.Release(auth, encodedSwap, r, vs, initiator, recipient)
// 	if err != nil {
// 		log.Printf("releaseTx error: %s\n", err)
// 		return err
// 	}

// 	fmt.Printf("releaseTx sent: %s", releaseTx.Hash().Hex())

// 	//send executeSwap tx to origin chain

// 	auth, err = getAuth(originClient, fromAddress, privateKey)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	executeSwapTx, err := originInstance.DirectExecuteSwap(auth, encodedSwap, r, vs, initiator, recipient)
// 	if err != nil {
// 		log.Printf("executeSwap error: %s\n", err)
// 		return err
// 	}

// 	fmt.Printf("executeSwap sent: %s", executeSwapTx.Hash().Hex())

// 	return nil
// }

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main2() {
	// callBridge()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
