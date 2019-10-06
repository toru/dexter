package index

import (
	"fmt"
)

type Issuer struct {
	hashAlgo uint8
}

// NewIssuer returns a new ID Issuer
func NewIssuer(algo string) (Issuer, error) {
	isr := Issuer{}
	if !IsSupported(algo) {
		return isr, fmt.Errorf("unsupported hash algorithm")
	}
	isr.hashAlgo = hashAlgoDict[algo]
	return isr, nil
}
