/*
 * WIP: SHA-1 based object addresser.
 */

package index

import (
	"crypto/sha1"
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
