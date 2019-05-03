package index

import (
	"bytes"
	"testing"
)

func TestNewIDFromHexDigest(t *testing.T) {
	orig := NewDexIDFromString("dexter")
	digest := DexIDToHexDigest(orig)

	repro, err := NewDexIDFromHexDigest(digest)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(repro[:], orig[:]) {
		t.Errorf("Got: %x, Want: %x", repro, orig)
	}
}
