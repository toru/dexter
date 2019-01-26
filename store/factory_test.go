package store

import (
	"testing"
)

func TestGetStore(t *testing.T) {
	t.Run("with in-memory store", func(t *testing.T) {
		expectedName := "Memory Store"

		s, ok := GetStore("memory")
		if !ok {
			t.Error("Got: true, Want: false")
		}
		if name := s.Name(); name != expectedName {
			t.Errorf("Got: %s, Want: %s", name, expectedName)
		}
	})
}
