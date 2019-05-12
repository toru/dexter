// Package index holds various indexing related data structures and algorithms.
package index

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

const (
	DexHexIDLen = 56
)

// DexID is a 224-bit (or 28-byte) long primary key that every
// object inside Dexter has been alloted with.
type DexID = [sha256.Size224]byte

// NewDexIDFromString returns a new DexID based on the given string.
func NewDexIDFromString(src string) DexID {
	return sha256.Sum224([]byte(src))
}

// NewDexIDFromHexDigest returns a new DexID based on the given hex digest.
func NewDexIDFromHexDigest(src string) (DexID, error) {
	if len(src) != DexHexIDLen {
		return DexID{}, errors.New("invalid dexter id")
	}
	rv := DexID{}
	raw, err := hex.DecodeString(src)
	if err != nil {
		return rv, err
	}
	copy(rv[:], raw)
	return rv, nil
}

// DexIDToHexDigest returns the hexadecimal representation of the given
// DexID as a string. Sadly, DexID can't be used as a method receiver.
func DexIDToHexDigest(id DexID) string {
	return hex.EncodeToString(id[:])
}