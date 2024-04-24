package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

type Message2 struct {
	EncodedSwap string
	Sign        string
	Recipient   string
	Initiator   string
}

func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func AirdropCPToken(w http.ResponseWriter, r *http.Request) {
	claimRet := TxStatus{
		Status:  1,
		TokenId: 0,
		TxHash:  "",
		Err:     ""}
	q := r.URL.Query()
	tokenAddress := common.HexToAddress(q.Get("token"))
	toAddress := common.HexToAddress(q.Get("to"))
	amount := q.Get("amount")
	fmt.Printf("tokenAddress:%s toAddress:%s amount:%s\n", tokenAddress, toAddress, amount)
	targetClient, err := ethclient.Dial("https://rpc.devnet.onenesslabs.io")
	if err != nil {
		log.Println(err)

		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	targetInstance, err := NewCPToken(tokenAddress, targetClient)

	privateKey, err := crypto.HexToECDSA(goDotEnvVariable("PRIVATE_KEY"))
	if err != nil {
		log.Println(err)
		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")

		claimRet.Err = "cannot assert type: publicKey is not of type *ecdsa.PublicKey"
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return

	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("wallet address:%s\n", fromAddress)

	auth, err := getAuth(targetClient, fromAddress, privateKey)
	if err != nil {
		log.Println(err)

		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}
	amountWei := ToWei(amount, 18)
	mintTx, err := targetInstance.Mint(auth, toAddress, amountWei)
	if err != nil {
		log.Printf("mint error: %s\n", err)
		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	log.Printf("mint %s", mintTx.Hash().Hex())

	claimRet.Status = 0
	claimRet.TxHash = mintTx.Hash().Hex()
	claimRet.TokenId = 0
	jsonStr, err := json.Marshal(claimRet)
	fmt.Fprintf(w, string(jsonStr))
	return
}

func GetReceipt(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	txHashArg := q.Get("txhash")
	targetClient, err := ethclient.Dial("https://rpc.devnet.onenesslabs.io")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}
	_txHash := common.HexToHash(txHashArg)
	receipt, err := targetClient.TransactionReceipt(context.Background(), _txHash)

	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	contractAbi, err := abi.JSON(strings.NewReader(string(SBTERC721ABI)))
	if err != nil {
		log.Fatal(err)
	}
	// bigInt := new(big.Int)
	var tokenId = new(big.Int)
	for _, vLog := range receipt.Logs {
		// log.Printf("%"vLog)
		log.Printf("event SigHash:%s %s", vLog.Topics[0].Hex(), logTransferSigHash.Hex())
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			var transferEvent EventTransfer
			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.from = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.to = common.HexToAddress(vLog.Topics[2].Hex())
			// log.Printf(string(vLog.Topics[3].Hex()))
			_, isSuccess := tokenId.SetString(string(vLog.Topics[3].Hex())[2:], 16)
			if isSuccess {
				log.Printf(tokenId.String())
			}

			// tokenId = common.S256(transferEvent.TokenId)
			fmt.Println(transferEvent.from.Hex()) // foo
			fmt.Println(transferEvent.to.Hex())
			// fmt.Println(transferEvent.TokenId.Uint64())
		}
	}

	fmt.Fprintf(w, tokenId.String())

}

type TxStatus struct {
	// 0. SUCCESS 1.ERROR 2.PENDING
	Status  int    `json:"status"`
	TokenId uint   `json:"tokenId"`
	TxHash  string `json:"txHash"`
	Err     string `json:"err"`
}

func ClaimSBT(w http.ResponseWriter, r *http.Request) {
	claimRet := TxStatus{
		Status:  1,
		TokenId: 0,
		TxHash:  "",
		Err:     ""}
	q := r.URL.Query()
	tokenAddress := common.HexToAddress(q.Get("token"))
	toAddress := common.HexToAddress(q.Get("to"))
	fmt.Printf("claimSBT tokenAddress:%s toAddress:%s", tokenAddress, toAddress)
	targetClient, err := ethclient.Dial("https://rpc.devnet.onenesslabs.io")
	if err != nil {
		log.Println(err)

		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	targetInstance, err := NewSBTERC721(tokenAddress, targetClient)

	privateKey, err := crypto.HexToECDSA(goDotEnvVariable("PRIVATE_KEY"))
	if err != nil {
		log.Println(err)
		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("claimSBT cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		claimRet.Err = "claimSBT cannot assert type: publicKey is not of type *ecdsa.PublicKey"
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))

		return

	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("claimSBT wallet address:%s\n", fromAddress)

	auth, err := getAuth(targetClient, fromAddress, privateKey)
	if err != nil {
		log.Println(err)
		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}
	// amountWei := ToWei(amount, 18)
	mintTx, err := targetInstance.Mint(auth, toAddress)
	if err != nil {
		log.Printf("claimSBT error: %s\n", err)
		claimRet.Err = err.Error()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	log.Printf("claimSBT %s", mintTx.Hash().Hex())
	AttemptLimit := 3

	var receipt *types.Receipt

	isSuccess := false
	for i := 1; i <= AttemptLimit; i++ {
		time.Sleep(time.Second * 2)
		receipt, err = targetClient.TransactionReceipt(context.Background(), mintTx.Hash())
		if err != nil {
			log.Printf("receipt %s", err)

		} else {
			isSuccess = true
			break
		}
	}

	if !isSuccess {
		claimRet.Err = "txreceipt not found"
		claimRet.Status = 2
		claimRet.TxHash = mintTx.Hash().Hex()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	contractAbi, err := abi.JSON(strings.NewReader(string(SBTERC721ABI)))
	if err != nil {
		claimRet.Err = err.Error()
		claimRet.Status = 1
		claimRet.TxHash = mintTx.Hash().Hex()
		jsonStr, _ := json.Marshal(claimRet)
		fmt.Fprintf(w, string(jsonStr))
		return
	}

	var tokenId = new(big.Int)
	for _, vLog := range receipt.Logs {
		// log.Printf("%"vLog)
		log.Printf("event SigHash:%s %s", vLog.Topics[0].Hex(), logTransferSigHash.Hex())
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			var transferEvent EventTransfer
			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				claimRet.Err = err.Error()
				claimRet.Status = 1
				claimRet.TxHash = mintTx.Hash().Hex()
				jsonStr, _ := json.Marshal(claimRet)
				fmt.Fprintf(w, string(jsonStr))
				return
			}

			transferEvent.from = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.to = common.HexToAddress(vLog.Topics[2].Hex())
			_, isSuccess = tokenId.SetString(vLog.Topics[3].Hex()[2:], 16)
			if !isSuccess {
				claimRet.Err = "parse topic error"
				claimRet.Status = 2
				claimRet.TxHash = mintTx.Hash().Hex()
				jsonStr, _ := json.Marshal(claimRet)
				fmt.Fprintf(w, string(jsonStr))
				return
			}
			// tokenId = transferEvent.tokenId
			log.Printf("transferEvent from:%s to:%s tokenId:%s\n", transferEvent.from.Hex(), transferEvent.to.Hex(), tokenId.String())
		}
	}

	claimRet.Status = 0
	claimRet.TxHash = mintTx.Hash().Hex()
	claimRet.TokenId = uint(tokenId.Int64())
	jsonStr, err := json.Marshal(claimRet)
	// log.Printf("err:%s json:%s", err, jsonStr)
	fmt.Fprintf(w, string(jsonStr))
	return
}

type EventTransfer struct {
	from    common.Address
	to      common.Address
	TokenId string
}

func main() {
	// callBridge()

	mux := http.NewServeMux()
	mux.HandleFunc("/airdrop/cp", AirdropCPToken)
	mux.HandleFunc("/airdrop/sbt", ClaimSBT)
	mux.HandleFunc("/receipt/get", GetReceipt)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
