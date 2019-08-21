package feed

import (
	"bytes"
	"encoding/xml"
)

const atomHint = "http://www.w3.org/2005/Atom"
const searchLimit = 1024

// FeedFormat attempts to detects the feed format of the given byte slice.
// Detection is performed on a best-effort basis.
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
	if upper > searchLimit {
		upper = searchLimit
	}
	return bytes.Contains(doc[:upper], []byte(atomHint))
}
