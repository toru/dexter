package storage

import (
	"sync"

	"github.com/toru/dexter/feed"
	"github.com/toru/dexter/index"
	"github.com/toru/dexter/subscription"
)

// MemoryStore is a simple memory-backed storage engine.
type MemoryStore struct {
	subsMux       sync.RWMutex
	subscriptions map[string]subscription.Subscription

	feedsMux sync.RWMutex
	feeds    map[string]feed.Feed
	entries  map[string]feed.Entry
}

// NewMemoryStore returns a new MemoryStore.
func NewMemoryStore() (*MemoryStore, error) {
	ret := &MemoryStore{}
	ret.subscriptions = make(map[string]subscription.Subscription)
	ret.feeds = make(map[string]feed.Feed)
	ret.entries = make(map[string]feed.Entry)
	return ret, nil
}

// Name returns the name of the storage engine.
func (s MemoryStore) Name() string {
	return "Memory Store"
}

// Subscriptions returns a slice of stored subscriptions.
func (s *MemoryStore) Subscriptions() []subscription.Subscription {
	s.subsMux.RLock()
	defer s.subsMux.RUnlock()

	subs := make([]subscription.Subscription, 0, len(s.subscriptions))
	for _, sub := range s.subscriptions {
		subs = append(subs, sub)
	}
	return subs
}

// WriteSubscription stores the given subscription.
func (s *MemoryStore) WriteSubscription(sub *subscription.Subscription) error {
	s.subsMux.Lock()
	defer s.subsMux.Unlock()

	s.subscriptions[sub.ID.String()] = *sub
	return nil
}

// NumSubscriptions returns the number of stored subscriptions.
func (s *MemoryStore) NumSubscriptions() int {
	s.subsMux.RLock()
	defer s.subsMux.RUnlock()

	return len(s.subscriptions)
}

// Feeds returns a slice of stored feeds.
func (s *MemoryStore) Feeds() []feed.Feed {
	s.feedsMux.RLock()
	defer s.feedsMux.RUnlock()

	feeds := make([]feed.Feed, 0, s.NumSubscriptions())
	for _, f := range s.feeds {
		feeds = append(feeds, f)
	}
	return feeds
}

// Feed finds a feed by its unique ID.
func (s *MemoryStore) Feed(id index.ID) (feed.Feed, bool) {
	s.feedsMux.RLock()
	defer s.feedsMux.RUnlock()

	f, ok := s.feeds[id.String()]
	if !ok {
		return nil, false
	}
	return f, true
}

// WriteFeed stores the given feed, indexed by its ID.
func (s *MemoryStore) WriteFeed(f feed.Feed) error {
	s.feedsMux.Lock()
	defer s.feedsMux.Unlock()

	s.feeds[f.SubscriptionID().String()] = f
	return nil
}

// Entries returns a slice of feed.Entry that belongs to a feed that
// matches the given identifier.
func (s *MemoryStore) Entries(id index.ID) []feed.Entry {
	s.feedsMux.RLock()
	defer s.feedsMux.RUnlock()

	var rv []feed.Entry
	f, ok := s.feeds[id.String()]
	if ok {
		rv = f.Entries()
	}
	return rv
}
