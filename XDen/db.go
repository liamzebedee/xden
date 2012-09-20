package XDen

import (
	"github.com/liamzebedee/go-qrp"
)

type DB struct {
	localNode qrp.Node // QRP node used for RPC
	localContact Contact
	
	timeout int
}

func NewDB() (*DB) {
	
	return nil
}