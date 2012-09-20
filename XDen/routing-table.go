package XDen

import (
	"container/list"
	"sync"
)

/*
DEN implements a different bucket storage method. 
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
	
}

func (*routingTable) addContact() () {
	
}


type bucket struct {
	maxSize int
	list.List
	mutex sync.Mutex
}

func NewBucket(maxSize int32) bucket {
	bucket := bucket {}
	// Init maxSize
	// Init list
	// Init mutex
}

func (bucket *bucket) Add(contact *Contact) (error) {
	
	if bucket.Len() == bucket.maxSize {
		// Max capacity
		
	} else {
		// Add node to end of bucket
		// TODO
	}
	return nil
}