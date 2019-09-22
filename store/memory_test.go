package store

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

	sub, _ := subscription.New("https://ep.torumk.com/feed")
	if err := s.WriteSubscription(sub); err != nil {
		t.Error(err)
	}
	if n := s.NumSubscriptions(); n != 1 {
		t.Errorf("Got: %d, Want: 1", n)
	}
}

func TestWriteFeed(t *testing.T) {
	s, _ := NewMemoryStore()
	k := index.NewSHA224DexIDFromString("ok")
	f := feed.NewAtomFeed()
	f.SetSubscriptionID(k.Value())

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
