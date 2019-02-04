package store

import (
	"testing"
)

func TestCreateSubscription(t *testing.T) {
	s, _ := NewMemoryStore()
	if n := s.NumSubscriptions(); n != 0 {
		t.Errorf("Got: %d, Want: 0", n)
	}
}
