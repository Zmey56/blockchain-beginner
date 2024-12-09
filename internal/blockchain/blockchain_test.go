package blockchain

import "testing"

func TestCalculateHash(t *testing.T) {
	block := Block{Index: 1, Timestamp: "2024-11-28", Data: "test", PrevHash: "123"}
	expectedHash := CalculateHash(block)
	if expectedHash == "" {
		t.Errorf("Expected hash to be non-empty")
	}
}

func TestGenerateBlock(t *testing.T) {
	prevBlock := Block{Index: 0, Hash: "123"}
	newBlock := GenerateBlock(prevBlock, "test", 4)
	if newBlock.Index != 1 {
		t.Errorf("Expected index to be 1, got %d", newBlock.Index)
	}
	if newBlock.PrevHash != "123" {
		t.Errorf("Expected prevHash to be 123, got %s", newBlock.PrevHash)
	}
}
