package fake

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateEthAddress() string {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return crypto.Keccak256Hash(bytes).Hex()
}
