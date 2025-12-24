package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/text"
	"math/big"
	"os"
	"time"
)

func CreateAccount() {
	//Create account (wallet) begin
	account := solana.NewWallet()
	fmt.Println("account private key:", account.PrivateKey)
	fmt.Println("account public key:", account.PublicKey())

	//account private key: 2RMPqxZ8R5RJQXmFuAHU4hHNQD7h8whKCXWXWDHQhEP53YL1mPrPYEwt3RoQFmQiGQdKL7aP3YanAaLA5L155i6E
	//account public key: YiL8hfnMAjQmXLH97gVsqGijDw8rHY97FApCcqfKZH4
	// Create a new RPC client:
	client := rpc.New(rpc.TestNet_RPC)

	// Airdrop 1 SOL to the new account:
	out, err := client.RequestAirdrop(
		context.TODO(),
		account.PublicKey(),
		solana.LAMPORTS_PER_SOL*1,
		rpc.CommitmentFinalized,
	)

	if err != nil {
		panic(err)
	}
	fmt.Println("airdrop transaction signature:", out)
	//Create account (wallet) end
}

func Transfer() {

	//Transfer Sol from one wallet to another wallet

	rpcClient := rpc.New(rpc.DevNet_RPC)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		panic(err)
	}

	// Load the account that you will send funds FROM:
	accountFrom, err := solana.PrivateKeyFromSolanaKeygenFileBytes([]byte(`2RMPqxZ8R5RJQXmFuAHU4hHNQD7h8whKCXWXWDHQhEP53YL1mPrPYEwt3RoQFmQiGQdKL7aP3YanAaLA5L155i6E`))
	if err != nil {
		panic(err)
	}
	fmt.Println("accountFrom private key:", accountFrom)
	fmt.Println("accountFrom public key:", accountFrom.PublicKey())

	// The public key of the account that you will send sol TO:
	accountTo := solana.MustPublicKeyFromBase58("TODO")
	// The amount to send (in lamports);
	// 1 sol = 1000000000 lamports
	amount := uint64(3333)

	if true {
		// Airdrop 1 sol to the account so it will have something to transfer:
		out, err := rpcClient.RequestAirdrop(
			context.TODO(),
			accountFrom.PublicKey(),
			solana.LAMPORTS_PER_SOL*1,
			rpc.CommitmentFinalized,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("airdrop transaction signature:", out)
		time.Sleep(time.Second * 5)
	}
	//---------------

	recent, err := rpcClient.GetLatestBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amount,
				accountFrom.PublicKey(),
				accountTo,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(accountFrom.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if accountFrom.PublicKey().Equals(key) {
				return &accountFrom
			}
			return nil
		},
	)
	if err != nil {
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}
	spew.Dump(tx)
	// Pretty print the transaction:
	tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	// Send transaction, and wait for confirmation:
	sig, err := confirm.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(sig)

}

// 获取最新区块
func GetLatestBlockhash() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	example, err := client.GetLatestBlockhash(
		context.Background(),
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(example)
	spew.Dump(example.Value)
}

// // 查询账户余额
func GetBalance() {
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)

	pubKey := solana.MustPublicKeyFromBase58("7xLk17EQQ5KLDLDe44wCmupJKJjTGd8hs3eSVVhCx932")
	out, err := client.GetBalance(
		context.TODO(),
		pubKey,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
	spew.Dump(out.Value) // total lamports on the account; 1 sol = 1000000000 lamports

	var lamportsOnAccount = new(big.Float).SetUint64(uint64(out.Value))
	// Convert lamports to sol:
	var solBalance = new(big.Float).Quo(lamportsOnAccount, new(big.Float).SetUint64(solana.LAMPORTS_PER_SOL))

	// WARNING: this is not a precise conversion.
	fmt.Println("◎", solBalance.Text('f', 10))
}

// 实时交易订阅
func SignatureSubscribe() {
	ctx := context.Background()
	fmt.Println("url", rpc.TestNet_WS)
	client, err := ws.Connect(context.Background(), rpc.TestNet_WS)
	if err != nil {
		panic(err)
	}

	txSig := solana.MustSignatureFromBase58("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	sub, err := client.SignatureSubscribe(
		txSig,
		"",
	)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		got, err := sub.Recv(ctx)
		if err != nil {
			panic(err)
		}
		spew.Dump(got)
	}
}

func main() {
	//GetLatestBlockhash()
	//GetBalance()
	//CreateAccount()
	//Transfer()
	SignatureSubscribe()
}
