package feed

import (
	"crypto/sha256"
)

type Feed interface {
	ID() string
	Title() string
	SubscriptionID() [sha256.Size224]byte
}
