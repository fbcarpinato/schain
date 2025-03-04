package main

import (
	"fmt"

	"github.com/fbcarpinato/schain/internal/blockchain"
)

func main() {
	blockchain := blockchain.NewBlockchain()

	blockchain.AddBlock("Test first block")

	for _, block := range blockchain.Blocks {
		fmt.Printf("Block %x with data %s\n", block.Hash, block.Data)
	}
}
