package feed

import (
	"testing"
)

var rssDoc = []byte(`<?xml version="1.0" encoding="UTF-8"?><rss version="2.0"></rss>`)
var atomDoc = []byte(`<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom"></feed>`)

func TestFeedFormat(t *testing.T) {
	if got := FeedFormat(atomDoc); got != AtomFeedFormat {
		t.Errorf("Got: %d, Want: %d", got, AtomFeedFormat)
	}
	if got := FeedFormat(rssDoc); got != RSS2FeedFormat {
		t.Errorf("Got: %d, Want: %d", got, UnknownFeedFormat)
	}
}
