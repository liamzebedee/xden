package XDen

import (
	"bytes"
)

type DB struct {
	localContact *Contact
}

func NewDB() (*DB) {
	
	return nil
}

// Gets an object addressed by key
func (db *DB) Get(key Key) (data bytes.Buffer) {
	// TODO
	return *bytes.NewBufferString("")
}

// Puts an object into the network
// If it already exists, we try to put the diff of the current object and the new object
func (db *DB) Put(key Key, data bytes.Buffer) {
	
}

func (db *DB) Remove(key Key) {
	
}