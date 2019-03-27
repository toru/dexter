// Package subscription implements subscription management.
package subscription

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/toru/dexter/feed"
	"github.com/toru/dexter/index"
)

// Subscription represents a subscription to a data feed.
type Subscription struct {
	ID      index.DexID // Unique ID
	FeedURL url.URL     // URL of the data endpooint

	unreachable  bool // Consider using a enum
	checksum     index.DexID
	createdAt    time.Time
	lastSyncedAt time.Time
}

// New returns a new Subscription.
func New(feedURL string) (*Subscription, error) {
	u, err := url.Parse(feedURL)
	if err != nil {
		return nil, err
	}

	s := &Subscription{
		ID:      index.NewDexIDFromString(feedURL),
		FeedURL: *u,
	}
	return s, nil
}

// IsOffline returns a boolean indicating the data feed reachability.
func (s *Subscription) IsOffline() bool {
	// TODO(toru): Somehow allow to retry. Maybe exponential backoff.
	return s.unreachable
}

// FeedSync downloads the data feed, parses it, and returns a Feed.
func (s *Subscription) FeedSync() (feed.Feed, error) {
	if len(s.FeedURL.String()) == 0 {
		return nil, fmt.Errorf("subscription has no FeedURL")
	}
	if s.unreachable {
		return nil, fmt.Errorf("unreachable: %s", s.FeedURL.String())
	}

	// TODO(toru): Proper HTTP client with network timeout.
	resp, err := http.Get(s.FeedURL.String())
	s.lastSyncedAt = time.Now().UTC()
	if err != nil {
		s.unreachable = true
		return nil, err
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		s.unreachable = true
		return nil, fmt.Errorf("sync failure (%d): %s", resp.StatusCode, s.FeedURL.String())
	}

	checksum := sha256.Sum224(payload)
	if bytes.Equal(s.checksum[:], checksum[:]) {
		return nil, fmt.Errorf("no new content: %s", s.FeedURL.String())
	}
	s.checksum = checksum

	if feed.IsAtomFeed(payload) {
		af, err := feed.ParseAtomFeed(payload)
		if err != nil {
			return nil, err
		}
		af.SetSubscriptionID(s.ID)
		return af, nil
	} else {
		return nil, fmt.Errorf("unknown syndication format")
	}
}
