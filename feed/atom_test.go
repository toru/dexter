package feed

import (
	"testing"
)

func TestIsAtomFeed(t *testing.T) {
	t.Run("with an atom document", func(t *testing.T) {
		atomDoc := []byte(`<feed xmlns="http://www.w3.org/2005/Atom"></feed>`)
		if !IsAtomFeed(atomDoc) {
			t.Errorf("Got: false, Want: true")
		}
	})

	t.Run("with an rss document", func(t *testing.T) {
		rssDoc := []byte(`<?xml version="1.0" encoding="UTF-8" ?><rss version="2.0"></rss>`)
		if IsAtomFeed(rssDoc) {
			t.Errorf("Got: true, Want: false")
		}
	})
}
