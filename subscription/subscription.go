// Package subscription implements subscription management.
package subscription

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Subscription represents a subscription to a data feed.
type Subscription struct {
	feedURL      url.URL
	unreachable  bool // Consider using a enum
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

	if resp.StatusCode != 200 {
		s.unreachable = true
		return fmt.Errorf("sync failure (%d): %s", resp.StatusCode, s.feedURL.String())
	}

	// TODO(toru): Parse the response.
	return nil
}
