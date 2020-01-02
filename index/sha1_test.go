package index

import (
	"testing"
)

func TestSHA1SetValueFromString(t *testing.T) {
	id := &SHA1DexID{}
	id.SetValueFromString("evergreen")
	want := "59acab5d0a5ee07d665440a21ec311c06bcdfe93"
	got := id.HexString()
	if got != want {
		t.Errorf("Got: %s, Want: %s", got, want)
	}
}

func TestSHA1SetValueFromHexString(t *testing.T) {
	id := &SHA1DexID{}
	src := "invalid"
	if err := id.SetValueFromHexString(src); err == nil {
		t.Error("Got: ok, Want: fail")
	}

	src = "53e71790fd8b19a7e7299454ab0b98ac96b33295"
	if err := id.SetValueFromHexString(src); err != nil {
		t.Error(err)
	}
}
