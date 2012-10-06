package XDen

/*

*/

import (
	"bytes"
)

type DB struct {
}

func NewDB() (*DB) {
	// TODO
	return nil
}

// Gets an object addressed by key
func (db *DB) Get(key Key) (data bytes.Buffer) {
	// TODO
	return *bytes.NewBufferString("")
}



func (db *DB) Remove(key Key) {
	// TODO
}