package feed

import (
	"crypto/sha256"
)

// Feed is a common interface for all data feed formats.
type Feed interface {
	ID() string
	Title() string
	SubscriptionID() [sha256.Size224]byte
}
