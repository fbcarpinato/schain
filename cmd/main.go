package main

import (
	"fmt"

	"github.com/fbcarpinato/schain/internal/blockchain"
)

func main() {
	block := blockchain.NewBlock("Test data", []byte{})

	fmt.Printf("Block hash: %x\n", block.Hash)
}
