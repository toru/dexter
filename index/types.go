// Package index holds various indexing related data structures and algorithms.
package index

import (
	"crypto/sha256"
)

// DexID is a 224-bit (or 28-byte) long primary key that every
// object inside Dexter has been alloted with.
type DexID = [sha256.Size224]byte
