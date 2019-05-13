package feed

import (
	"bytes"
)

const atomHint = "http://www.w3.org/2005/Atom"
const detectionLimit = 1024

const (
	UnknownFeedFormat = iota
	RSS1FeedFormat
	RSS2FeedFormat
	AtomFeedFormat
)

func FeedFormat(doc []byte) int {
	if isAtomFeed(doc) {
		return AtomFeedFormat
	}
	return UnknownFeedFormat
}

// Heuristically determines if the document is an Atom feed by searching
// for the format namespace. Search is given up after 1024 bytes.
func isAtomFeed(doc []byte) bool {
	upper := len(doc)
	if upper > detectionLimit {
		upper = detectionLimit
	}

	// TODO: Sadly this is not enough, as shown by the NASA feeds.
	// Partially parse the given document and evaluate from there.
	return bytes.Contains(doc[:upper], []byte(atomHint))
}
