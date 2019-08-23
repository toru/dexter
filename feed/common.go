package feed

import (
	"github.com/toru/dexter/index"
)

const (
	UnknownFeedFormat = iota
	RSS1FeedFormat
	RSS2FeedFormat
	AtomFeedFormat
)

// Feed is a common interface for all data feed formats.
type Feed interface {
	ID() string
	SubscriptionID() index.DexID
	Title() string
	Format() uint
	Entries() []Entry
	SetSubscriptionID(id index.DexID)
}

// Entry is a common interface for all data entry formats.
type Entry interface {
	FeedID() index.DexID
	ID() string
	Summary() string
}
