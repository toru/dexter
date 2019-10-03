package index

import (
	"fmt"
)

type Issuer struct {
}

func NewIssuer(algo string) (Issuer, error) {
	isr := Issuer{}
	if !IsSupported(algo) {
		return isr, fmt.Errorf("unsupported hash algorithm")
	}
	return isr, nil
}
