// Package subscription implements subscription management.
package subscription

import (
	"net/url"
	"time"
)

// Subscription represents a subscription to a data feed.
type Subscription struct {
	feedURL      url.URL
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
