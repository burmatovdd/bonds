package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashPassword(password string) string {
	data := []byte(password)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
