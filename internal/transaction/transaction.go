package transaction

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"strconv"

	"github.com/fbcarpinato/schain/internal/wallet"
)

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    int
	Signature []byte
}

func NewTransaction(sender *wallet.Wallet, receiver *wallet.Wallet, amount int) (*Transaction, error) {
	if sender.Balance < amount {
		return nil, errors.New("insufficient sender balance")
	}

	tx := &Transaction{
		Sender:   sender.GetAddress(),
		Receiver: receiver.GetAddress(),
		Amount:   amount,
	}

	amountStr := strconv.Itoa(amount)
	hash := sha256.Sum256(bytes.Join([][]byte{
		[]byte(tx.Sender),
		[]byte(tx.Receiver),
		[]byte(amountStr),
	}, nil))

	r, s, err := ecdsa.Sign(rand.Reader, sender.PrivateKey, hash[:])
	if err != nil {
		return nil, err
	}

	rBytes := r.Bytes()
	sBytes := s.Bytes()
	signature := make([]byte, 64)
	copy(signature[32-len(rBytes):32], rBytes)
	copy(signature[64-len(sBytes):64], sBytes)

	tx.Signature = signature

	sender.Balance -= amount
	receiver.Balance += amount

	return tx, nil
}
