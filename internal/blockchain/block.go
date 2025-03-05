package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"

	"github.com/fbcarpinato/schain/internal/transaction"
)

const targetPrefix = "000"

type Block struct {
	Timestamp    int64
	PreviousHash []byte
	Hash         []byte
	Transactions []*transaction.Transaction
	Nonce        int
}

func NewBlock(transactions []*transaction.Transaction, previousHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		PreviousHash: previousHash,
		Transactions: transactions,
		Nonce:        0,
	}

	block.Mine()

	return block
}

func (b *Block) IsMined() bool {
	return bytes.HasPrefix(b.Hash, []byte(targetPrefix))
}

func (b *Block) Mine() {
	for {
		b.Hash = b.CalculateHash()
		if b.IsMined() {
			break
		}
		b.Nonce++
	}
}

func (b *Block) CalculateHash() []byte {
	var txHashes [][]byte
	for _, tx := range b.Transactions {
		txData := bytes.Join([][]byte{
			[]byte(tx.Sender),
			[]byte(tx.Receiver),
			[]byte(strconv.Itoa(tx.Amount)),
			tx.Signature,
		}, []byte{})
		txHashes = append(txHashes, txData)
	}

	data := bytes.Join([][]byte{
		[]byte(strconv.FormatInt(b.Timestamp, 10)),
		b.PreviousHash,
		bytes.Join(txHashes, []byte{}),
		[]byte(strconv.Itoa(b.Nonce)),
	}, []byte{})

	hash := sha256.Sum256(data)
	return hash[:]
}
