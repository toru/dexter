package store

import (
	"sync"

	"github.com/toru/dexter/subscription"
)

// MemoryStore is a simple memory-backed storage engine.
type MemoryStore struct {
	subsMux       sync.RWMutex
	subscriptions map[string]subscription.Subscription
}

// NewMemoryStore returns a new MemoryStore.
func NewMemoryStore() (*MemoryStore, error) {
	ret := &MemoryStore{}
	ret.subscriptions = make(map[string]subscription.Subscription)
	return ret, nil
}

// Name returns the name of the storage engine.
func (s MemoryStore) Name() string {
	return "Memory Store"
}

// Subscriptions returns a slice of stored subscriptions.
func (s *MemoryStore) Subscriptions() []subscription.Subscription {
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

	// TODO(toru): Precalculate a sha224 fingerprint
	s.subscriptions[sub.FeedURL.String()] = *sub
	return nil
}

// NumSubscriptions returns the number of stored subscriptions.
func (s *MemoryStore) NumSubscriptions() int {
	s.subsMux.RLock()
	defer s.subsMux.RUnlock()

	return len(s.subscriptions)
}
