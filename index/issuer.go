package index

import (
	"fmt"
)

type Issuer struct {
	hashAlgo uint8
}

// NewIssuer returns a new ID Issuer.
func NewIssuer(algo string) *Issuer {
	return &Issuer{}
}

// SetAlgo sets the hashing algorithm.
func (i *Issuer) SetAlgo(algo string) error {
	if !IsSupported(algo) {
		return fmt.Errorf("unsupported hash algorithm: %s", algo)
	}
	i.hashAlgo = hashAlgoDict[algo]
	return nil
}
