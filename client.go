package main

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
	err = callBridge(m)
	if err != nil {
		fmt.Fprintf(w, "fail")
	} else {
		fmt.Fprintf(w, "success")
	}

}

func callBridge(m Message) error {
	targetMesonAddr := common.HexToAddress("0x244Bb9e8D1d1b2B73182D594B87691eD8708105e")

	targetClient, err := ethclient.Dial("https://rpc.devnet.onenesslabs.io")
	if err != nil {
		log.Println(err)
		return err
	}

	originMesonAddr := common.HexToAddress("0x0d12d15b26a32e72A3330B2ac9016A22b1410CB6")
	originClient, err := ethclient.Dial("https://sepolia.drpc.org")
	if err != nil {
		log.Println(err)
		return err
	}

	_ = targetClient

	privateKey, err := crypto.HexToECDSA(goDotEnvVariable("PRIVATE_KEY"))
	if err != nil {
		log.Println(err)
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

	//send tx to origin chain

	nonce, err := originClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return err
	}

	gasPrice, err := originClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	originInstance, err := NewMeson(originMesonAddr, originClient)

	fmt.Println(originInstance, err)

	releaseTx, err := originInstance.DirectRelease(auth, encodedSwap, r, vs, recipient, initiator)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("releaseTx sent: %s", releaseTx.Hash().Hex())

	//send tx to target chain
	targetNonce, err := targetClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return err
	}

	targetGasPrice, err := targetClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	auth = bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(targetNonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = targetGasPrice

	targetInstance, err := NewMeson(targetMesonAddr, targetClient)
	swaptx, err := targetInstance.ExecuteSwap(auth, encodedSwap, r, vs, recipient, true)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("swaptx sent: %s", swaptx.Hash().Hex())
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

func main() {
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
