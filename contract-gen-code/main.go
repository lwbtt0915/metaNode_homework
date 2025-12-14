package main

import (
	"context"
	token "contract-gen-code/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	contractAddr = "0x00267645F5677350F740da9C8586Ded87816D811"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/aca471a53c104ef2b9182b5cf7637f45")

	if err != nil {
		log.Fatal(err)
	}

	simpleCounterContract, err := token.NewToken(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f58221423788589ed88f53eab6ea7499c3a503eb9de42b8b89bf9ad98a3e8bbe")
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(100000))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := simpleCounterContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := simpleCounterContract.GetCount(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is value saving in contract equals to origin value:", valueInContract)
}
