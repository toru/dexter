package store

import (
	"testing"

	"github.com/toru/dexter/subscription"
)

func TestCreateSubscription(t *testing.T) {
	s, _ := NewMemoryStore()
	if n := s.NumSubscriptions(); n != 0 {
		t.Errorf("Got: %d, Want: 0", n)
	}

	sub := subscription.New()
	if err := s.CreateSubscription(sub); err != nil {
		t.Error(err)
	}
	if err := s.CreateSubscription(sub); err == nil {
		t.Error("Got: success, Want: failure")
	}
	if n := s.NumSubscriptions(); n != 1 {
		t.Errorf("Got: %d, Want: 1", n)
	}
}
