package store

import (
	"testing"
)

func TestGetStore(t *testing.T) {
	if _, ok := GetStore("memory"); ok {
		t.Errorf("Got: true, Want: false")
	}
}
