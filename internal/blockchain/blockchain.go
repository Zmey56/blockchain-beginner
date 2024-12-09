package blockchain

import "time"

// GenerateBlock creates a new block based on the previous block and the data provided
func GenerateBlock(prevBlock Block, transactions []Transaction, difficulty int) Block {
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
	}

	// Execute Proof-of-Work
	pow := NewProofOfWork(&newBlock, difficulty)
	hash, nonce := pow.Run()

	newBlock.Hash = hash
	newBlock.Nonce = nonce
	return newBlock
}
