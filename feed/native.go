package feed

import (
	"crypto/sha256"
)

type Feed interface {
	SubscriptionID() [sha256.Size224]byte
}
