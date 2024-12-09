package blockchain

type Transaction struct {
	ID       string              // Unique transaction ID
	Inputs   []TransactionInput  // Transaction Inputs
	Outputs  []TransactionOutput // Transaction outputs
	Coinbase bool                // Indicates that this is the initial transaction
}

type TransactionInput struct {
	PrevTxID  string // Hash of the previous transaction
	OutputIdx int    // Index of the output of the previous transaction
	Signature string // Signature confirming the right to use funds
}

type TransactionOutput struct {
	Value     int    // Transfer amount
	Recipient string // Recipient's address
}
