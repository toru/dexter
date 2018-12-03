package feed

import (
	"bytes"
	"encoding/xml"
	"time"
)

const atomHint = "http://www.w3.org/2005/Atom"
const detectionLimit = 1024

// AtomAuthor represents the "atom:author" element defined in
// RFC 4287 Section 4.2.1
type AtomAuthor struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
}

// AtomCategory represents the "atom:category" element defined in
// RFC 4287 Section 4.2.2
type AtomCategory struct {
	Term   string `xml:"term,attr"`
	Scheme string `xml:"scheme,attr"`
	Label  string `xml:"label,attr"`
}

// AtomGenerator represents the "atom:generator" element defined in
// RFC 4287 Section 4.2.4
type AtomGenerator struct {
	URI     string `xml:"uri,attr"`
	Version string `xml:"version,attr"`
	Value   string `xml:",chardata"`
}

// AtomFeed represents the top-level container element of the Atom
// feed document as defined in RFC 4287 Section 4.1.1
type AtomFeed struct {
	Author   AtomAuthor   `xml:"author"`
	Category AtomCategory `xml:"category"`
	// Contributor
	Generator AtomGenerator `xml:"generator"`
	Icon      string        `xml:"icon"`
	ID        string        `xml:"id"`
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
