package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// GenerateAddress Address generation based on random data
func GenerateAddress() string {
	data := time.Now().String()
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
