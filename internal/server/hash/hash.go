package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

type HashService struct {
	method Hash
}

type Hash interface {
	Password(password string) string
}

func (service *HashService) Password(password string) string {
	data := []byte(password)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
