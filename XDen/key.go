package XDen

import (
	"encoding/hex"
)

const (
	KEY_SIZE = 256 / 8
)

// Fundamental data type used for identifying nodes and objects
type Key [KEY_SIZE]byte

// XOR the target key against the current key
func (source *Key) Xor(target *Key) (result *Key) {
	for i := 0; i < KEY_SIZE; i++ {
		result[i] = source[i] ^ target[i]
	}
	return
}

// Returns the string representation of the key
func (key *Key) String() string {
	return hex.EncodeToString(key[0:KEY_SIZE])
}

// Equates this key to another key, returning true if they are equal
func (key *Key) Equals(other *Key) bool {
	for i := 0; i < KEY_SIZE; i++ {
		if key[i] != other[i] {
			return false
		}
	}
	return true
}

// 
func (key *Key) Less(other *Key) bool {
	for i := 0; i < keyLength; i++ {
		if key[i] != other[i] {
			return key[i] < other[i]
		}
	}
	return false
}

// Returns true if the nth digit is 1
func (key *Key) PrefixN(n int) bool {
	for i := 0; i < keyLength; i++ {
		for j := 0; j < 8; j++ {
			if (key[i] >> uint8(7-j)) & 0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return keyLength*8 - 1
}

// Returns the 
// Returns which bucket the node should be stored in (prefix length). 
// Determined by the number of leading 0 bits in the XOR of our node ID with the target node ID
func (key *Key) Prefix() (ret int) {
	// Iterates through each bit in the key, and when it reaches a set bit, it returns the current bitcount
	for i := 0; i < keyLength; i++ {
		for j := 0; j < 8; j++ {
			if (key[i] >> uint8(7-j)) & 0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return keyLength*8 - 1
}

// Decodes a Key from a Base16 string
func DecodeKey(data string) (result *Key) {
	decoded, _ := hex.DecodeString(data)
	for i := 0; i < keyLength; i++ {
		result[i] = decoded[i]
	}
	return
}