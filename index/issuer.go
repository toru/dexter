package index

import (
	"fmt"
	"sync"
)

type IssuerSingleton struct {
	hashAlgo uint8
}

var issuerInstance *IssuerSingleton
var issuerOnce sync.Once

// GetIssuer returns a new ID Issuer.
func GetIssuer() *IssuerSingleton {
	issuerOnce.Do(func() {
		issuerInstance = &IssuerSingleton{}
	})
	return issuerInstance
}

// SetAlgo sets the hashing algorithm.
func (i *IssuerSingleton) SetAlgo(algo string) error {
	if !IsSupported(algo) {
		return fmt.Errorf("unsupported hash algorithm: %s", algo)
	}
	i.hashAlgo = hashAlgoDict[algo]
	return nil
}

// CreateID returns a newly created ID
func (i *IssuerSingleton) CreateID() ID {
	if i.hashAlgo == DexIDTypeSHA1 {
		return &SHA1DexID{}
	}
	return &SHA224DexID{}
}
