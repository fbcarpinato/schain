package main

import (
	"fmt"

	"github.com/fbcarpinato/schain/internal/blockchain"
	"github.com/fbcarpinato/schain/internal/transaction"
	"github.com/fbcarpinato/schain/internal/wallet"
)

func main() {
	chain := blockchain.NewBlockchain()

	wallet1, _ := wallet.NewWallet()
	wallet2, _ := wallet.NewWallet()

	tx, err := transaction.NewTransaction(wallet1, wallet2, 10.0)
	if err != nil {
		fmt.Printf("Transaction creation failed: %v\n", err)
		return
	}

	chain.AddBlock([]*transaction.Transaction{tx})

	tx2, _ := transaction.NewTransaction(wallet2, wallet1, 5.0)
	chain.AddBlock([]*transaction.Transaction{tx2})

	for i, block := range chain.Blocks {
		fmt.Printf("\nBlock %d (%x):\n", i, block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PreviousHash)
		fmt.Printf("Transactions: %d\n", len(block.Transactions))

		for t, transaction := range block.Transactions {
			fmt.Printf("  Tx %d: %s => %s (%d coins)\n",
				t,
				transaction.Sender,
				transaction.Receiver,
				transaction.Amount)
		}
	}

	fmt.Printf("\nBlockchain valid: %t\n", chain.IsValid())

	fmt.Printf("\nWallet 1 balance: %d\n", wallet1.Balance)
	fmt.Printf("Wallet 2 balance: %d\n", wallet2.Balance)
}
