/*
 * WIP: SHA-1 based object addresser.
 */

package index

import (
	"crypto/sha1"
	"encoding/hex"
)

const (
	SHA1DexIDLen    = sha1.Size
	SHA1DexIDHexLen = sha1.Size * 2
)

type SHA1DexID struct {
	value [SHA1DexIDLen]byte
}

// Algo implements the ID interface.
func (id SHA1DexID) Algo() uint8 {
	return DexIDTypeSHA1
}

// Value implements the ID interface.
func (id SHA1DexID) Value() []byte {
	return id.value[:]
}

// HexString implements the ID interface.
func (id SHA1DexID) HexString() string {
	return hex.EncodeToString(id.value[:])
}
