package blockchain

import "bytes"

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", []byte{})

	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

func (bc *Blockchain) LatestBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *Blockchain) AddBlock(data string) {
	bc.Blocks = append(
		bc.Blocks,
		NewBlock(data, bc.LatestBlock().Hash),
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
