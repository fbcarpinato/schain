package blockchain

import (
	"bytes"

	"github.com/fbcarpinato/schain/internal/transaction"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock([]*transaction.Transaction{}, []byte{})

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func (bc *Blockchain) LatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(transactions []*transaction.Transaction) {
	bc.Blocks = append(
		bc.Blocks,
		NewBlock(transactions, bc.LatestBlock().Hash),
	)
}

func (bc *Blockchain) IsValid() bool {
	if len(bc.Blocks) == 0 {
		return true
	}

	for i := 1; i < len(bc.Blocks); i++ {
		prev, curr := bc.Blocks[i-1], bc.Blocks[i]

		if !bytes.Equal(curr.Hash, curr.CalculateHash()) ||
			!bytes.Equal(prev.Hash, prev.CalculateHash()) ||
			!bytes.Equal(curr.PreviousHash, prev.Hash) ||
			!curr.IsMined() || !prev.IsMined() {
			return false
		}
	}

	return true
}
