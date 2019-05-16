package feed

import (
	"bytes"
	"encoding/xml"
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
	if isRSS2Feed(doc) {
		return RSS2FeedFormat
	} else if isAtomFeed(doc) {
		return AtomFeedFormat
	}
	return UnknownFeedFormat
}

func isRSS2Feed(doc []byte) bool {
	rf := &struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
	}{}
	if err := xml.Unmarshal(doc, rf); err != nil {
		return false
	}
	return rf.Version == "2.0"
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
