package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

const targetPrefix = "000"

type Block struct {
	Timestamp    int64
	PreviousHash []byte
	Hash         []byte
	Data         string
	Nonce        int
}

func NewBlock(data string, previousHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousHash,
		Data:         data,
		Nonce:        0,
	}

	block.MineBlock()

	return block
}

func (b *Block) IsMined() bool {
	return bytes.HasPrefix(b.Hash, []byte(targetPrefix))
}

func (b *Block) MineBlock() {
	for {
		b.Hash = b.CalculateHash()

		if b.IsMined() {
			break
		}

		b.Nonce++
	}
}

func (b *Block) CalculateHash() []byte {
	data := bytes.Join([][]byte{
		[]byte(strconv.FormatInt(b.Timestamp, 10)),
		b.PreviousHash,
		[]byte(b.Data),
		[]byte(strconv.Itoa(b.Nonce)),
	}, []byte{})

	hash := sha256.Sum256(data)

	return hash[:]
}
