package p2p

import (
	"fmt"
	"net"
)

type Node struct {
	Address string
	Peers   map[string]net.Conn
}

func NewNode(address string) *Node {
	return &Node{
		Address: address,
		Peers:   map[string]net.Conn{},
	}
}

func (n *Node) Start() error {
	listener, err := net.Listen("tcp", n.Address)
	if err != nil {
		return err
	}
	defer listener.Close()
	defer fmt.Printf("Shutting down node on address %s", n.Address)

	fmt.Printf("Starting a new node on address %s", n.Address)

	for {
		_, err := listener.Accept()
		if err != nil {
			return err
		}
	}
}
