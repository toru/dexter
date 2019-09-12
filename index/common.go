// Package index holds various indexing related data structures and algorithms.
package index

const (
	DexIDTypeSHA1 = iota
	DexIDTypeSHA224
	DexIDTypeUUID
)

const (
	DexHexIDLen = 56
)

// ID is a common interface for all unique identifiers.
type ID interface {
	Algo() uint8
	Value() []byte
	String() string
	SetValue(val []byte)
}
