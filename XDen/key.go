package XDen

import (
	"encoding/hex"
	"bytes"
	"github.com/tkawachi/bitset"
)

// Fundamental data type used for identifying nodes and objects
type Key struct { bitset.BitSet }
const KEY_SIZE = 256

// Returns a new key of size KEY_SIZE
func NewKey() (Key) {
	return Key{*bitset.New(KEY_SIZE)}
}

// XOR the target key against the current key and return the result
func (source *Key) Xor(target *Key) (result Key) {
	return Key{*source.SymmetricDifference(&target.BitSet)}
}

// Returns a Hex (Base16) representation of the key as a string
func (key *Key) String() (string) {
	var buf bytes.Buffer
	key.WriteTo(&buf)
	return hex.EncodeToString(buf.Bytes())
}

func (key *Key) Equals(other *Key) bool {
	return (key.Equal(&other.BitSet))
}

func (key *Key) Less(other *Key) bool {
	// Dumping the string is more efficient than using BitSet.Set
	bits_1 := key.DumpAsBits()
	bits_2 := key.DumpAsBits()
	
	// Iterate through bits in string
	for i := 0; i < KEY_SIZE; i++ {
		if bits_1[i] != bits_2[i] {
			return bits_1[i] < bits_2[i]
		}
	}
	
	return false
}

func (key *Key) isSet(bit uint) bool {
	return key.Test(bit)
}

// Decodes a hexadecimal string into a Key
func (key *Key) DecodeKey(data string) (result *Key, err error) {
	decodedData, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(decodedData)
	// TODO Check if ReadFrom works correctly
	bitset, err := bitset.ReadFrom(buf)
	if err != nil {
		return nil, err
	}
	
	return &Key{*bitset}, nil
}

func (key *Key) Prefix() Prefix {
	return Prefix{key, 0}	
}

// An n-bit part of a key
type Prefix struct {
	*Key
	end uint // Indicates end of the prefix
}

func NewPrefix() (Prefix) {
	key := NewKey()
	return Prefix{ &key, 0 }	
}

// Returns true if target is in range of this prefix
func (prefix *Prefix) inRange(target *Prefix) bool {
	// Iterate through all bits in prefix
		// if prefix[bit] != target[prefix]
			// return false
	return true	
}