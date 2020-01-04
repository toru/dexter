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
	SetValueFromString(val string)
	SetValueFromHexString(val string) error
}

// IsSupported validates if the indexer supports the given hash algorithm.
func IsSupported(algo string) bool {
	algo = strings.ToLower(algo)
	return algo == "sha1" || algo == "sha224"
}
