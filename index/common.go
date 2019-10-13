// Package index holds various indexing related data structures and algorithms.
package index

import (
	"strings"
)

const (
	DexIDTypeSHA1 = iota
	DexIDTypeSHA224
)

var hashAlgoDict = map[string]uint8{
	"sha1":   DexIDTypeSHA1,
	"sha224": DexIDTypeSHA224,
}

// ID is a common interface for all unique identifiers.
type ID interface {
	Algo() uint8
	Value() []byte
	HexString() string
	SetValue(val []byte)
	SetValueFromHexString(val string) error
}

// IsSupported validates if the indexer supports the given hash algorithm.
// Dexter currently only supports SHA-224.
func IsSupported(algo string) bool {
	return strings.ToLower(algo) == "sha224"
}
