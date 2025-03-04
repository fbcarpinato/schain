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

	if len(bc.Blocks) == 1 {
		return bytes.Equal(bc.Blocks[0].Hash, bc.Blocks[0].CalculateHash())
	}

	for i := 1; i < len(bc.Blocks)-1; i++ {
		previousBlock := bc.Blocks[i-1]
		currentBlock := bc.Blocks[i]

		if bytes.Equal(currentBlock.Hash, currentBlock.CalculateHash()) {
			return false
		}

		if bytes.Equal(previousBlock.Hash, previousBlock.CalculateHash()) {
			return false
		}

		if bytes.Equal(currentBlock.PreviousHash, previousBlock.Hash) {
			return false
		}

		if !currentBlock.IsMined() || !previousBlock.IsMined() {
			return false
		}
	}

	return true
}
