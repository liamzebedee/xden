package XDen

import (
	"net"
	"github.com/liamzebedee/go-qrp"
)

// Holds info about a contact in the network
type Contact struct {
	nodeID Key
	addr   net.UDPAddr
}

// Call a procedure on this contact's node
func (contact *Contact) Call(method string, args, reply interface{}, db *DB) {
	// Call procedure with timeout
	db.localNode.Call(method, &contact.addr, args, reply, db.timeout)
	
	// Update node after response
	
}

// Returns if the contact is closer to 'target' than 'other'
func (contact *Contact) Closer(other *Contact, target *Key) (bool) {
	distance1 := *contact.nodeID.Xor(target)
	distance2 := *other.nodeID.Xor(target)
	return distance1.Less(&distance2)
}