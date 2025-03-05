package main

import (
	"fmt"

	"github.com/fbcarpinato/schain/internal/blockchain"
	"github.com/fbcarpinato/schain/internal/wallet"
)

func main() {
	blockchain := blockchain.NewBlockchain()

	blockchain.AddBlock("Test first block")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Block %x with data %s\n", block.Hash, block.Data)
	}

	fmt.Printf("Blockchain valid status: %t\n", blockchain.IsValid())

	wallet, err := wallet.NewWallet()

	if err != nil {
		fmt.Printf("Error while creating a new wallet: %s", err)
	}

	fmt.Printf("Created a new wallet with address: %s\n", wallet.GetAddress())
}
