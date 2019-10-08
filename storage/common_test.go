package storage

import (
	"testing"
)

func TestGetStore(t *testing.T) {
	t.Run("with in-memory store", func(t *testing.T) {
		cfg := Config{"memory"}
		expectedName := "Memory Store"

		s, err := GetStore(cfg)
		if err != nil {
			t.Error(err)
		}
		if name := s.Name(); name != expectedName {
			t.Errorf("Got: %s, Want: %s", name, expectedName)
		}
	})
	t.Run("with bogus store", func(t *testing.T) {
		cfg := Config{"bogus"}
		if _, err := GetStore(cfg); err == nil {
			t.Errorf("Got: nil, Want: error")
		}
	})
}
