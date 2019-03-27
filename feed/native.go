package feed

import (
	"github.com/toru/dexter/index"
)

// Feed is a common interface for all data feed formats.
type Feed interface {
	ID() string
	Title() string
	SubscriptionID() index.DexID
}
