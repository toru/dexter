package index

import (
	"bytes"
	"testing"
)

func TestSetValueFromHexString(t *testing.T) {
	id := &SHA224DexID{}
	src := "invalid"
	if err := id.SetValueFromHexString(src); err == nil {
		t.Error("Got: ok, Want: fail")
	}

	src = "ab66ac985b784cd6966b63122df2f73fc756aef8599530f9011a2c14"
	if err := id.SetValueFromHexString(src); err != nil {
		t.Error(err)
	}
}

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
