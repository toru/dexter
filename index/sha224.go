package index

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

const (
	SHA224DexIDLen    = sha256.Size224
	SHA224DexIDHexLen = sha256.Size224 * 2
)

// DexID is a 224-bit (or 28-byte) long primary key that every
// object inside Dexter has been alloted with.
type DexID = [sha256.Size224]byte

// ValidateHexID returns a boolean indicating the validity of the given
// hexadecimal string. Mostly syntax sugar at this point.
func ValidateHexID(digest string) bool {
	return len(digest) == SHA224DexIDHexLen
}

type SHA224DexID struct {
	value [sha256.Size224]byte
}

func NewSHA224DexIDFromString(src string) ID {
	return &SHA224DexID{sha256.Sum224([]byte(src))}
}

// Temporary constructor for ID migration work
func NewSHA224DexIDFromLegacyDexID(id DexID) ID {
	return &SHA224DexID{id}
}

// Algo implements the ID interface.
func (id SHA224DexID) Algo() uint8 {
	return DexIDTypeSHA224
}

// Value implements the ID interface.
func (id SHA224DexID) Value() []byte {
	return id.value[:]
}

// String implements the ID interface.
func (id SHA224DexID) String() string {
	return hex.EncodeToString(id.value[:])
}

// SetValue implements the ID interface.
func (id *SHA224DexID) SetValue(val []byte) {
	copy(id.value[:], val)
}

// SetValueFromHex implements the ID interface.
func (id *SHA224DexID) SetValueFromHexString(val string) error {
	if len(val) != SHA224DexIDHexLen {
		return errors.New("invalid hex string")
	}
	raw, err := hex.DecodeString(val)
	if err != nil {
		return err
	}
	copy(id.value[:], raw)
	return nil
}
