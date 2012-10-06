package XDen

import (
	"container/list"
)

/*
Like Kademlia, we start with a single bucket. When this bucket reaches it's maxSize, one of two things can happen:

1. The bucket is split
2. Old nodes are pinged (if we cannot split)

Splitting is an operation that occurs only if our own node ID falls within the range of the bucket.
The bucket being split is replaced by two new buckets each with half the range of the old bucket and 
the nodes from the old bucket are distributed among the two new ones.

Unlike Kademlia, XDen imposes different size limits on buckets.
The purpose of the routing table is to provide the most efficient lookup and storage mechanism for nodes, 
so the closest online nodes to a key can be found fast. 

The Kademliua routing table is made of 160 buckets and all buckets have a fixed size of 20 nodes. 
We implement a routing modification, NR128-A, described in 'Sub-Second Lookups on a Large-Scale Kademlia-Based Overlay' 
to improve lookup performance, by increasing bucket size. In this paper it was showed that:

a)	Performance improves logarithmically with bucket size
b)	Maintenance traffic (PING, FIND) grows linearly with bucket size

Given the structure of a Kademlia routing table, on average, the first bucket is used in half of the lookups, 
the second bucket in a quarter of the lookups, and so forth. We improve performance by resizing buckets according to 
their use: the first buckets hold 128, 64, 32, and 16 nodes respectively, while the rest of the bucket sizes remain at 8 nodes. 

Technically, this is implemented by sizing the root bucket to 128 nodes. When we split a bucket, the left/right bucket that isn't in the range
of our node inherits the size of it's parent. The other bucket is half the size of its parent, 8 nodes at minimum.

We also implement an optimization described in part of chapter 4.1 of the full Kademlia paper - a contact replacement cache.
*/

type routingTable struct {
	root *bucket // Bucket at the top of the tree
	replacementCache *bucket
	localContact *Contact
}

const ROOT_BUCKET_SIZE = 128
const CACHE_BUCKET_SIZE = 32

func NewRoutingTable() (*routingTable) {
	root := NewBucket(ROOT_BUCKET_SIZE, true, 0)
	cache := NewBucket(CACHE_BUCKET_SIZE, false, 0)
	r := routingTable{ root: root, replacementCache: cache }
	return &r
}

// Wrapper function for getBucketForKey
func (routingTable *routingTable) getBucketForContact(contact *Contact) (bucket *bucket) {
	return routingTable.getBucketForKey(contact.localID, routingTable.root)	
}

func (routingTable *routingTable) getBucketForKey(key Key, start *bucket) (bucket *bucket) {
	// Walk the tree
	if start.hasChildren() {
		// Traverse left (0) or right (1) depending on the value
		if key.Test(start.depth) {
			return routingTable.getBucketForKey(key, start.right)
		} else {
			return routingTable.getBucketForKey(key, start.left)
		}
	} else {
		// If we are at the bottom
		return start
	}
	return nil
}

// Adds a contact to the appropriate bucket in the routing table
// Will return false if the contact failed to be added
func (routingTable *routingTable) addContact(contact *Contact) (contactAdded bool) {
	appropriateBucket := routingTable.getBucketForContact(contact)
	
	if hasNode := appropriateBucket.update(contact); hasNode {
		return true
	}
	
	// Try to add contact
	if contactAdded := appropriateBucket.addContact(contact); !contactAdded {
		// Bucket at max size
		// Check if our own nodeID falls within the range of the bucket
		if(appropriateBucket.localNodeInRange) {
			leftBucketSize, rightBucketSize := appropriateBucket.maxSize, appropriateBucket.maxSize
			newDepth := appropriateBucket.depth + 1
			
			// If the new right bucket will be in range of our node	
			if rightBucketInRange := routingTable.localContact.nodeID.Test(newDepth); rightBucketInRange {
				// Create left bucket with same size as parent
				leftBucket := NewBucket(leftBucketSize, false, newDepth)
				
				if(appropriateBucket.maxSize < 32) {
					rightBucketSize = 16
				} else {
					rightBucketSize = appropriateBucket.maxSize / 2
				}
				
				// Create right bucket half the size of parent, at least 16
				rightBucket := NewBucket(rightBucketSize, true, newDepth)
				
				appropriateBucket.split(leftBucket, rightBucket)
			} else {
				// Opposite of above
				rightBucket := NewBucket(rightBucketSize, true, newDepth)
				
				if(appropriateBucket.maxSize < 32) {
					leftBucketSize = 16
				} else {
					leftBucketSize = appropriateBucket.maxSize / 2
				}
				
				leftBucket := NewBucket(leftBucketSize, false, newDepth)
				
				appropriateBucket.split(leftBucket, rightBucket)
			}
		} else {			
			leastRecentlySeen := appropriateBucket.Front().Value.(*Contact)
			if(leastRecentlySeen.IsOnline()) {
				appropriateBucket.update(leastRecentlySeen)
				routingTable.cacheContact(contact)
				return false
			} else {
				// Else replace them
				// TODO
				return true
			}
		}
	}
	
	return false
}

func (routingTable *routingTable) cacheContact(contact *Contact) (contactAdded bool) {
	return false // TODO
}

type bucket struct {
	left      		  *bucket // 0
	right     		  *bucket // 1
	maxSize   		  int	  // > 16
	depth     		  uint
	localNodeInRange  bool
					  
	 		          list.List // Front of list is least recently seen
}

func NewBucket(maxSize int, localNodeInRange bool, depth uint) (b *bucket) {
	b = &bucket{ maxSize: maxSize, localNodeInRange: localNodeInRange, depth: depth }
	b.List = *list.New()
	return b
}

func (bucket *bucket) hasChildren() bool {
	return ((bucket.left != nil) && (bucket.right != nil))
}

func (bucket *bucket) isFull() bool {
	return (bucket.Len() == bucket.maxSize)
}

// Updates a contact, moving it to the most-recently seen category.
// Returns false if it couldn't be updated (wasn't found)
func (bucket *bucket) update(contact *Contact) bool {
	for e := bucket.Front(); e != nil; e = e.Next() {
		if e.Value.(*Contact).nodeID.Equals(&contact.nodeID) {
			// If the node is already stored, we move it to the front of the list
			bucket.MoveToBack(e)
			return true
		}
	}
	
	return false
}

// Returns false if the contact wasn't added (if the bucket is at max size)
func (bucket *bucket) addContact(contact *Contact) bool {
	if bucket.isFull() {
		return false
	}
	bucket.PushBack(contact)
	return true
}

// Splits a bucket, allocating nodes the left and right buckets appropriately, and then sets it's children to those buckets
func (bucket *bucket) split(left, right *bucket) {
	bucket.left, bucket.right = left, right
	// Go through all nodes, reallocate
	for e := bucket.Front(); e != nil; e = e.Next() {
		contact := e.Value.(Contact)
		if contact.nodeID.isSet(right.depth) {
			
		}
	}
}