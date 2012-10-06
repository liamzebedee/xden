package XDen

import (
	"net"
	"github.com/liamzebedee/go-qrp"
)

const MAX_TIMEOUT = 1

// A contact is a collection of data associated with a node in a network
type Contact struct {
	node qrp.Node // QRP node used for RPC
	nodeID Key
	localID Key // nodeID XOR localNode.nodeID
	addr   net.UDPAddr
}

// Call a procedure on this contact's node
func (contact *Contact) Call(method string, args, reply interface{}, db *DB) {
	// TODO
	// Call procedure with timeout
	//db.localNode.Call(method, &contact.addr, args, reply, db.timeout)
	
	// Update node after response
	
}

// PINGs the contact, returning if they respond and are thus online
func (contact *Contact) IsOnline() bool {
	// TODO
	return false
}

// Returns if the contact is closer to 'target' than 'other'
func (contact *Contact) Closer(other *Contact, target *Key) (bool) {
	distance1 := contact.nodeID.Xor(target)
	distance2 := other.nodeID.Xor(target)
	return distance1.Less(&distance2)
}