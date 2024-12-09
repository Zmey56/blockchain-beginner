package storage

type Storage interface {
	SaveBlock(hash string, data []byte) error
	GetBlock(hash string) ([]byte, error)
	Iterate(fn func(hash string, data []byte) error) error
	Close() error
}
