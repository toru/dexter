package feed

import (
	"bytes"
	"encoding/xml"
	"time"
)

const atomHint = "http://www.w3.org/2005/Atom"
const detectionLimit = 1024

// AtomFeed represents the top-level container element of the Atom
// feed document as described in RFC 4287 Section 4.1.1
type AtomFeed struct {
	// Author
	// Category
	// Contributor
	// Generator
	Icon string `xml:"icon"`
	ID   string `xml:"id"`
	// Links
	Logo     string    `xml:"logo"`
	Rights   string    `xml:"rights"`
	Subtitle string    `xml:"subtitle"`
	Title    string    `xml:"title"`
	Updated  time.Time `xml:"updated"`
	// Extentions
	// Entries
}

// Heuristically determines if the document is an Atom feed by searching
// for the format namespace. Search is given up after 1024 bytes.
func IsAtomFeed(doc []byte) bool {
	upper := len(doc)
	if upper > detectionLimit {
		upper = detectionLimit
	}
	return bytes.Contains(doc[:upper], []byte(atomHint))
}

// ParseAtomFeed parses the given byte slice as an AtomFeed.
func ParseAtomFeed(doc []byte) (*AtomFeed, error) {
	feed := &AtomFeed{}
	if err := xml.Unmarshal(doc, feed); err != nil {
		return nil, err
	}
	return feed, nil
}
