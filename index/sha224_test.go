package index

import (
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
