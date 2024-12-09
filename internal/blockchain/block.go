package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// Block represents a block in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
	Transactions []Transaction
}

// CalculateHash calculates the hash of a block
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}
