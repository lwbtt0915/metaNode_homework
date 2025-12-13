package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	fmt.Println("hello world.")

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/aca471a53c104ef2b9182b5cf7637f45")

	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(9829075)
	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	//查询区块信息
	fmt.Println(block.Number().Uint64())     //9829075
	fmt.Println(block.Time())                //1765599864
	fmt.Println(block.Difficulty().Uint64()) //0
	fmt.Println(block.Hash().Hex())          //0xd4f188919aa1a04c033fc8e9e493a8030228c6215b1fe049755d0cd22acaded7
	fmt.Println(len(block.Transactions()))   //139

	//发送交易
	privateKey, err := crypto.HexToECDSA("f58221423788589ed88f53eab6ea7499c3a503eb9de42b8b89bf9ad98a3e8bbe")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(100000) // in wei (1 eth)
	gasLimit := uint64(21000)   // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("bdc6cf849028fe95be67f0de7514512ef83488830786ab72f06759d0134eaf99")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
