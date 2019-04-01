// Package index holds various indexing related data structures and algorithms.
package index

import (
	"crypto/sha256"
	"encoding/hex"
)

// DexID is a 224-bit (or 28-byte) long primary key that every
// object inside Dexter has been alloted with.
type DexID = [sha256.Size224]byte

// NewDexIDFromString returns a new DexID based on the given string.
func NewDexIDFromString(src string) DexID {
	return sha256.Sum224([]byte(src))
}

// DexIDToHexDigest returns the hexadecimal representation of the given
// DexID as a string. Sadly, DexID can't be used as a method receiver.
func DexIDToHexDigest(id DexID) string {
	return hex.EncodeToString(id[:])
}
