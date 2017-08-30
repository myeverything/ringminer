package db

// ldb与table公用同一个interface
type Database interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Close()
	NewBatch() Batch
}

// batch&tablebatch公用同一个interface
type Batch interface {
	Put(key, value []byte) error
	Write() error
}