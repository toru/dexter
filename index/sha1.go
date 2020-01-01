/*
 * WIP: SHA-1 based object addresser.
 */

package index

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
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

// SetValue implements the ID interface.
func (id *SHA1DexID) SetValue(val []byte) {
	copy(id.value[:], val)
}

// SetValueFromString implements the ID interface.
func (id *SHA1DexID) SetValueFromString(val string) {
	id.value = sha1.Sum([]byte(val))
}

// SetValueFromHex implements the ID interface.
func (id *SHA1DexID) SetValueFromHexString(val string) error {
	if len(val) != SHA1DexIDHexLen {
		return errors.New("invalid hex string")
	}
	raw, err := hex.DecodeString(val)
	if err != nil {
		return err
	}
	copy(id.value[:], raw)
	return nil
}
