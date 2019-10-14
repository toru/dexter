package index

import (
	"testing"
)

func TestSetValueFromString(t *testing.T) {
	id := &SHA224DexID{}
	id.SetValueFromString("evergreen")
	want := "ca29aa62fbeb0497539a64be6d97f18b4b393773a98901513c20eb82"
	got := id.HexString()
	if got != want {
		t.Errorf("Got: %s, Want: %s", got, want)
	}
}

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
