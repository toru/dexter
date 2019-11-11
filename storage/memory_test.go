package storage

import (
	"testing"

	"github.com/toru/dexter/feed"
	"github.com/toru/dexter/index"
	"github.com/toru/dexter/subscription"
)

func TestWriteSubscription(t *testing.T) {
	s, _ := NewMemoryStore()
	if n := s.NumSubscriptions(); n != 0 {
		t.Errorf("Got: %d, Want: 0", n)
	}

	endpoint := "https://ep.torumk.com/feed"
	sub := subscription.New()
	sub.Init(endpoint)
	if err := s.WriteSubscription(sub); err != nil {
		t.Error(err)
	}
	if n := s.NumSubscriptions(); n != 1 {
		t.Errorf("Got: %d, Want: 1", n)
	}
}

func TestWriteFeed(t *testing.T) {
	s, _ := NewMemoryStore()
	k := &index.SHA224DexID{}
	k.SetValueFromString("ok")
	f := feed.NewAtomFeed()
	f.SetSubscriptionID(k)

	if err := s.WriteFeed(f); err != nil {
		t.Error(err)
	}
	if n := len(s.Feeds()); n != 1 {
		t.Errorf("Got: %d, Want: 1", n)
	}
	_, ok := s.Feed(k)
	if !ok {
		t.Errorf("Got: !ok, Want: ok")
	}
}
