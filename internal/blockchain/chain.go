package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Zmey56/blockchain-beginner/internal/storage"
)

type Blockchain struct {
	Blocks  []Block
	Storage storage.Storage
}

// AddBlock Adding a block with saving
func (bc *Blockchain) AddBlock(block Block) error {
	bc.Blocks = append(bc.Blocks, block)

	// Saving a block to the storage
	blockData, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return bc.Storage.SaveBlock(block.Hash, blockData)
}

// LoadBlockchain Loading the blockchain from the storage
func (bc *Blockchain) LoadBlockchain() error {
	err := bc.Storage.Iterate(func(hash string, data []byte) error {
		var block Block
		if err := json.Unmarshal(data, &block); err != nil {
			return err
		}
		bc.Blocks = append(bc.Blocks, block)
		return nil
	})
	if err != nil {
		return err
	}
	log.Println("The blockchain has been successfully loaded from the storage")
	return nil
}

func (bc *Blockchain) GetBalance(address string) int {
	balance := 0
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			for _, out := range tx.Outputs {
				if out.Recipient == address {
					balance += out.Value
				}
			}
		}
	}
	return balance
}

func CreateTransaction(sender, recipient string, amount int, bc *Blockchain) (*Transaction, error) {
	// Checking the sender's balance
	balance := bc.GetBalance(sender)
	if balance < amount {
		return nil, fmt.Errorf("недостаточно средств")
	}

	// Creating a transaction
	tx := &Transaction{
		Inputs: []TransactionInput{
			{PrevTxID: "", OutputIdx: 0, Signature: sender}, // Простая подпись
		},
		Outputs: []TransactionOutput{
			{Value: amount, Recipient: recipient},
			{Value: balance - amount, Recipient: sender}, // Сдача
		},
	}

	// Generating a unique transaction ID
	tx.ID = calculateTransactionID(tx)
	return tx, nil
}

func calculateTransactionID(tx *Transaction) string {
	data := fmt.Sprintf("%v", tx)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// CreateCoinbaseTransaction Creating a Coinbase transaction
func CreateCoinbaseTransaction(recipient string, amount int) *Transaction {
	tx := &Transaction{
		Inputs:   []TransactionInput{}, // У Coinbase-транзакций нет входов
		Outputs:  []TransactionOutput{{Value: amount, Recipient: recipient}},
		Coinbase: true,
	}
	tx.ID = calculateTransactionID(tx)
	return tx
}
