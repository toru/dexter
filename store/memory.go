package store

import (
	"fmt"
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

// CreateSubscription stores the given subscription.
func (s *MemoryStore) CreateSubscription(sub *subscription.Subscription) error {
	s.subsMux.Lock()
	defer s.subsMux.Unlock()

	// TODO(toru): Precalculate a sha224 fingerprint
	k := sub.FeedURL.String()
	if _, ok := s.subscriptions[k]; ok {
		return fmt.Errorf("duplicate key: %s", k)
	}
	s.subscriptions[k] = *sub
	return nil
}

// NumSubscriptions returns the number of stored subscriptions.
func (s *MemoryStore) NumSubscriptions() int {
	s.subsMux.RLock()
	defer s.subsMux.RUnlock()

	return len(s.subscriptions)
}
