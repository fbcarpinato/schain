package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	PreviousHash []byte
	Hash         []byte
	Data         string
}

func NewBlock(data string, previousHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousHash,
		Data:         data,
	}

	block.Hash = block.calculateHash()

	return block
}

func (b *Block) calculateHash() []byte {
	data := bytes.Join([][]byte{
		[]byte(strconv.FormatInt(b.Timestamp, 10)),
		b.PreviousHash,
		[]byte(b.Data),
	}, []byte{})

	hash := sha256.Sum256(data)

	return hash[:]
}
