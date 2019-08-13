package feed

import (
	"github.com/toru/dexter/index"
)

// Feed is a common interface for all data feed formats.
type Feed interface {
	ID() string
	Title() string
	SubscriptionID() index.DexID
	Entries() []Entry
	SetSubscriptionID(id index.DexID)
}

// Entry is a common interface for all data entry formats.
type Entry interface {
	FeedID() index.DexID
	ID() string
	Summary() string
}
