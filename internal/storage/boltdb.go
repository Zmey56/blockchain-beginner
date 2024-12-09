package storage

import (
	"fmt"

	"go.etcd.io/bbolt"
)

const blocksBucket = "blocks"

type BoltDB struct {
	db *bbolt.DB
}

// NewBoltDB Creating a new BoltDB storage
func NewBoltDB(path string) (*BoltDB, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(blocksBucket))
		return err
	})
	if err != nil {
		return nil, err
	}

	return &BoltDB{db: db}, nil
}

// SaveBlock Saving a block
func (b *BoltDB) SaveBlock(hash string, data []byte) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		return bucket.Put([]byte(hash), data)
	})
}

// GetBlock Getting a block
func (b *BoltDB) GetBlock(hash string) ([]byte, error) {
	var data []byte
	err := b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		if data == nil {
			return fmt.Errorf("block not found")
		}
		return nil
	})
	return data, err
}

// Close Closing the database
func (b *BoltDB) Close() error {
	return b.db.Close()
}

// Iterate Iterating over the blocks
func (b *BoltDB) Iterate(callback func(hash string, data []byte) error) error {
	return b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		return bucket.ForEach(func(k, v []byte) error {
			return callback(string(k), v)
		})
	})
}
