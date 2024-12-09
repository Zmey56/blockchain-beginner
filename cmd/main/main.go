package main

import (
	"fmt"
	"log"

	"github.com/Zmey56/blockchain-beginner/internal/blockchain"
	"github.com/Zmey56/blockchain-beginner/internal/storage"
	"github.com/Zmey56/blockchain-beginner/internal/utils"
)

func main() {
	// Open the storage and load the blockchain
	db, err := storage.NewBoltDB("blockchain.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	bc := &blockchain.Blockchain{Storage: db}
	if err := bc.LoadBlockchain(); err != nil {
		log.Fatalf("Error loading blockchain: %v", err)
	}

	// CLI interface
	var command string
	for {
		fmt.Println("Enter a command (gen-address, balance, send, view, exit):")
		fmt.Scanln(&command)

		switch command {
		case "gen-address":
			address := utils.GenerateAddress()
			fmt.Println("Your new address:", address)

		case "fund":
			var recipient string
			var amount int
			fmt.Println("Enter the recipient's address:")
			fmt.Scanln(&recipient)
			fmt.Println("Enter amount:")
			fmt.Scanln(&amount)

			// Creating a Coinbase Transaction
			coinbaseTx := blockchain.CreateCoinbaseTransaction(recipient, amount)

			// Adding a transaction to a new block
			newBlock := blockchain.GenerateBlock(bc.Blocks[len(bc.Blocks)-1], []blockchain.Transaction{*coinbaseTx}, 4)
			if err := bc.AddBlock(newBlock); err != nil {
				log.Printf("Error adding a block: %v", err)
			} else {
				fmt.Println("The balance has been replenished!")
			}

		case "balance":
			var address string
			fmt.Println("Enter the address:")
			fmt.Scanln(&address)
			fmt.Printf("Balance of address %s: %d\n", address, bc.GetBalance(address))

		case "send":
			var sender, recipient string
			var amount int
			fmt.Println("Enter the sender's address:")
			fmt.Scanln(&sender)
			fmt.Println("Enter the recipient's address:")
			fmt.Scanln(&recipient)
			fmt.Println("Enter the amount:")
			fmt.Scanln(&amount)

			tx, err := blockchain.CreateTransaction(sender, recipient, amount, bc)
			if err != nil {
				fmt.Println("Error creating transaction:", err)
				continue
			}

			// Add the transaction to a new block
			newBlock := blockchain.GenerateBlock(bc.Blocks[len(bc.Blocks)-1], []blockchain.Transaction{*tx}, 4)
			if err := bc.AddBlock(newBlock); err != nil {
				fmt.Printf("Error adding block: %v\n", err)
			} else {
				fmt.Println("Transaction added to block!")
			}

		case "view":
			for _, block := range bc.Blocks {
				fmt.Printf("Block %d: %s\n", block.Index, block.Hash)
				for _, tx := range block.Transactions {
					fmt.Printf("  Transaction %s: %d -> %s\n", tx.ID, tx.Outputs[0].Value, tx.Outputs[0].Recipient)
				}
			}

		case "exit":
			return

		default:
			fmt.Println("Unknown command.")
		}
	}
}
