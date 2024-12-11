package blockchain

import "testing"

func TestGenerateBlock(t *testing.T) {
	// Создаём предыдущий блок (генезис-блок)
	prevBlock := Block{
		Index:     0,
		Timestamp: "2024-11-28",
		Transactions: []Transaction{
			{
				ID: "genesis_tx",
				Outputs: []TransactionOutput{
					{Value: 100, Recipient: "address1"},
				},
			},
		},
		Hash: "1234567890abcdef",
	}

	// Creating a new block with a transaction
	newBlock := GenerateBlock(prevBlock, []Transaction{
		{
			ID: "tx1",
			Outputs: []TransactionOutput{
				{Value: 50, Recipient: "address2"},
			},
		},
	}, 2)

	// Checking the correctness of the index of the new block
	if newBlock.Index != prevBlock.Index+1 {
		t.Errorf("Index %d was expected, %d was received", prevBlock.Index+1, newBlock.Index)
	}

	// Checking that the hash of the previous block is correct
	if newBlock.PrevHash != prevBlock.Hash {
		t.Errorf("Hash %s was expected, %s was received", prevBlock.Hash, newBlock.PrevHash)
	}

	// Checking the transaction have been added
	if len(newBlock.Transactions) != 1 {
		t.Errorf("1 transaction was expected, %d was received", len(newBlock.Transactions))
	}

	// Checking the transaction data
	if newBlock.Transactions[0].Outputs[0].Recipient != "address2" {
		t.Errorf("Recipient address2 was expected, %s was received", newBlock.Transactions[0].Outputs[0].Recipient)
	}
}

func TestCreateTransaction(t *testing.T) {
	bc := Blockchain{}

	// Genesis block with initial funds
	genesisBlock := GenerateBlock(Block{}, []Transaction{
		{
			Outputs: []TransactionOutput{
				{Value: 100, Recipient: "address1"},
			},
		},
	}, 2)
	bc.Blocks = append(bc.Blocks, genesisBlock)

	// Creating a transaction
	tx, err := CreateTransaction("address1", "address2", 50, &bc)
	if err != nil {
		t.Fatalf("Transaction creation error: %v", err)
	}

	// Checking the transfer amount
	if tx.Outputs[0].Value != 50 {
		t.Errorf("50 expected, received %d", tx.Outputs[0].Value)
	}

	// Checking the delivery amount
	if tx.Outputs[1].Value != 50 {
		t.Errorf("Expected 50 (delivery), received %d", tx.Outputs[1].Value)
	}
}
