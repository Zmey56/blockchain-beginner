package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type ProofOfWork struct {
	Block  *Block
	Target string // Цель (количество начальных нулей в хеше)
}

// NewProofOfWork Create new Proof-of-Work
func NewProofOfWork(block *Block, difficulty int) *ProofOfWork {
	// The goal is set by a string of a given number of zeros
	target := strings.Repeat("0", difficulty)
	return &ProofOfWork{Block: block, Target: target}
}

// Run Performing calculations for PoW
func (pow *ProofOfWork) Run() (string, int) {
	var nonce int
	var hash string

	fmt.Println("Starting mining...")
	for {
		// Forming a string based on the block data and the current nonce
		data := pow.Block.Data + strconv.Itoa(pow.Block.Index) + pow.Block.PrevHash + strconv.Itoa(nonce)
		hashBytes := sha256.Sum256([]byte(data))
		hash = hex.EncodeToString(hashBytes[:])

		// Checking whether the hash satisfies the goal
		if strings.HasPrefix(hash, pow.Target) {
			break
		}
		nonce++
	}
	fmt.Printf("Mining is completed! Nonce: %d, Hash: %s\n", nonce, hash)
	return hash, nonce
}
