package main

import (
	"fmt"

	"github.com/fbcarpinato/schain/internal/p2p"
)

func main() {
	node := p2p.NewNode("localhost:8888")

	err := node.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}
