package storage

import (
	"testing"
)

func TestBoltDB(t *testing.T) {
	db, err := NewBoltDB("test.db")
	if err != nil {
		t.Fatalf("Database creation error: %v", err)
	}
	defer db.Close()

	data := []byte("test data")
	err = db.SaveBlock("test_hash", data)
	if err != nil {
		t.Fatalf("Block saving error: %v", err)
	}

	loadedData, err := db.GetBlock("test_hash")
	if err != nil {
		t.Fatalf("Block loading error: %v", err)
	}

	if string(loadedData) != string(data) {
		t.Errorf("%s data expected, %s received", string(data), string(loadedData))
	}
}
