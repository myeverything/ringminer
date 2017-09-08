package lrcdb

// interface for ldb and table
type Database interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Close()
	NewBatch() Batch
}

// interface for batch and table batch
type Batch interface {
	Put(key, value []byte) error
	Write() error
}