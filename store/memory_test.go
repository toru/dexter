package store

import (
	"testing"

	"github.com/toru/dexter/subscription"
)

func TestWriteSubscription(t *testing.T) {
	s, _ := NewMemoryStore()
	if n := s.NumSubscriptions(); n != 0 {
		t.Errorf("Got: %d, Want: 0", n)
	}

	sub := subscription.New()
	if err := s.WriteSubscription(sub); err != nil {
		t.Error(err)
	}
	if n := s.NumSubscriptions(); n != 1 {
		t.Errorf("Got: %d, Want: 1", n)
	}
}
