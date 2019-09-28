package storage

import (
	"testing"
)

func TestGetStore(t *testing.T) {
	t.Run("with in-memory store", func(t *testing.T) {
		expectedName := "Memory Store"

		s, err := GetStore("memory")
		if err != nil {
			t.Error(err)
		}
		if name := s.Name(); name != expectedName {
			t.Errorf("Got: %s, Want: %s", name, expectedName)
		}
	})
	t.Run("with bogus store", func(t *testing.T) {
		if _, err := GetStore("bogus"); err == nil {
			t.Errorf("Got: nil, Want: error")
		}
	})
}
