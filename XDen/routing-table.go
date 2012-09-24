package XDen

import (
	"container/list"
	"sync"
)

/*
XDen implements a different bucket storage method. 
Like the BitTorrent Mainline DHT, we start with a single bucket. When this bucket reaches it's
maxSize, one of two things can happen:

1. The bucket is split
2. Old nodes are pinged (if we cannot split)

Splitting is an operation that occurs only if our own node ID falls within the range of the bucket. 
The bucket being split is replaced by two new buckets each with half the range of the old bucket and 
the nodes from the old bucket are distributed among the two new ones.

Unlike Mainline DHT, DENS imposes different size limits on buckets
*/
type routingTable struct {
	root *bucket // Bucket at the top of the tree, *king of the jungle*
	replacementCache *bucket
}

func (routingTable *routingTable) getBucketForContact(contact *Contact) *bucket {
	// While the bucket still has children, 
	 
}

// Adds a contact to the appropriate bucket in the routing table
// Will return false if the contact failed to be added
func (routingTable *routingTable) addContact(contact *Contact) (contactAdded bool) {
	// Find appropriate bucket for contact
	bucket := routingTable.getBucketForContact(contact)
	
	for e := bucket.Front(); e != nil; e = e.Next() {
		if e.Value.(*Contact).nodeID.Equals(&contact.nodeID) {
			// If the node is already stored, we move it to the front of the list
			bucket.MoveToBack(e)
			return true
		}
	}
	
	// Check if the bucket is full
	if bucket.Len() == bucket.maxSize {	
		// Check if our own nodeID falls within the range of the bucket
			// Split bucket
			// Reallocate nodes
			// Add node
		// Else, PING the last seen node
		
			// If they reply, keep them
				// Add contact to replacement cache
			// Else replace them
	}
	return false
}


type bucket struct {
	left *bucket
	right *bucket
	maxSize int
	list.List
	mutex sync.Mutex
}

func (bucket *bucket) hasChildren() bool {
	return ((bucket.left != nil) && (bucket.right != nil))
}