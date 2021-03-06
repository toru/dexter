package storage

import (
	"errors"

	"github.com/toru/dexter/feed"
	"github.com/toru/dexter/index"
	"github.com/toru/dexter/subscription"
)

// Store is an interface that Storage Engines must implement.
type Store interface {
	Name() string
	Subscriptions() []subscription.Subscription
	WriteSubscription(sub *subscription.Subscription) error
	NumSubscriptions() int
	Feeds() []feed.Feed
	Feed(index.ID) (feed.Feed, bool)
	WriteFeed(f feed.Feed) error
	Entries(index.ID) []feed.Entry
}

// Config holds Storage Engine related settings.
type Config struct {
	Engine string // Name of the storage engine
}

// GetStore returns a Storage Engine based on the given name.
func GetStore(cfg Config) (Store, error) {
	switch cfg.Engine {
	case "memory":
		s, err := NewMemoryStore()
		if err != nil {
			return nil, err
		}
		return s, nil
	case "mysql", "mariadb":
		return nil, errors.New("work in progress")
	default:
		return nil, errors.New("unknown storage engine")
	}
}
