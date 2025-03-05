package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"github.com/btcsuite/btcutil/base58"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Balance    int
}

func NewWallet() (*Wallet, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return nil, errors.New("Failed to generate a private key")
	}

	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Balance:    100,
	}, nil
}

func (w *Wallet) GetAddress() string {
	firstHash := sha256.Sum256(w.PublicKey)
	secondHash := sha256.Sum256(firstHash[:])

	versionedPayload := append([]byte{0x00}, secondHash[:]...)

	checksum := sha256.Sum256(versionedPayload)
	checksum = sha256.Sum256(checksum[:])

	fullPayload := append(versionedPayload, checksum[:4]...)

	address := base58.Encode(fullPayload)

	return address

}
