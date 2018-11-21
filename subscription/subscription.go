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
)

// Subscription represents a subscription to a data feed.
type Subscription struct {
	feedURL      url.URL
	unreachable  bool // Consider using a enum
	checksum     [sha256.Size224]byte
	createdAt    time.Time
	lastSyncedAt time.Time
}

// New returns a new Subscription.
func New() *Subscription {
	return &Subscription{}
}

// SetFeedURL sets the URL for subscription.
func (s *Subscription) SetFeedURL(feedURL string) error {
	u, err := url.Parse(feedURL)
	if err != nil {
		return err
	}
	s.feedURL = *u
	return nil
}

// Sync downloads the data feed and parses it.
func (s *Subscription) Sync() error {
	if len(s.feedURL.String()) == 0 {
		return fmt.Errorf("subscription has no feedURL")
	}
	if s.unreachable {
		return fmt.Errorf("subscription is unreachable")
	}

	// TODO(toru): This is only for dev-purpose. Craft a proper HTTP
	// client with defensive settings like network timeout.
	resp, err := http.Get(s.feedURL.String())
	s.lastSyncedAt = time.Now().UTC()
	if err != nil {
		s.unreachable = true
		return err
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		s.unreachable = true
		return fmt.Errorf("sync failure (%d): %s", resp.StatusCode, s.feedURL.String())
	}

	checksum := sha256.Sum224(payload)
	if bytes.Equal(s.checksum[:], checksum[:]) {
		// There's no new content to process.
		return nil
	}
	s.checksum = checksum

	if feed.IsAtomFeed(payload) {
		// TODO(toru): Parse the response.
		fmt.Println("this is an atom feed")
	} else {
		return fmt.Errorf("unknown syndication format")
	}

	return nil
}
