package feed

import (
	"bytes"
	"encoding/xml"
	"time"

	"github.com/toru/dexter/index"
)

const atomHint = "http://www.w3.org/2005/Atom"
const detectionLimit = 1024

// atomText represents atom(Plain|XHTML)TextConstruct defined in
// RFC 4287 Section 3.1
type atomText struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",innerxml"`
}

// atomPerson represents the Person construct defined in
// RFC 4287 Section 3.2
type atomPerson struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
}

// AtomAuthor represents the "atom:author" element defined in
// RFC 4287 Section 4.2.1
type AtomAuthor = atomPerson

// AtomContributor represents the "atom:contributor" element defined in
// RFC 4287 Section 4.2.3
type AtomContributor = atomPerson

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

// AtomLink represents the "atom:link" element defined in
// RFC 4287 Section 4.2.7
type AtomLink struct {
	HRef     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr"`
	HRefLang string `xml:"hreflang,attr"`
	Type     string `xml:"type,attr"`
	Title    string `xml:"title,attr"`
	Length   string `xml:"length,attr"`
}

// AtomFeed represents the top-level container element of the Atom
// feed document as defined in RFC 4287 Section 4.1.1
type AtomFeed struct {
	Authors      []AtomAuthor      `xml:"author"`
	Categories   []AtomCategory    `xml:"category"`
	Contributors []AtomContributor `xml:"contributor"`
	Generator    AtomGenerator     `xml:"generator"`
	Icon         string            `xml:"icon"`
	ID_          string            `xml:"id"`
	Links        []AtomLink        `xml:"link"`
	Logo         string            `xml:"logo"`
	Rights       string            `xml:"rights"`
	Subtitle     string            `xml:"subtitle"`
	Title_       string            `xml:"title"`
	Updated      time.Time         `xml:"updated"`
	Entries      []AtomEntry       `xml:"entry"`

	// Dexter specific attributes
	subscriptionID index.DexID
}

// AtomEntry represents the "atom:entry" container element defined in
// RFC 4287 Section 4.1.2
type AtomEntry struct {
	Authors      []AtomAuthor      `xml:"author"`
	Categories   []AtomCategory    `xml:"category"`
	Content      atomText          `xml:"content"`
	Contributors []AtomContributor `xml:"contributor"`
	ID           string            `xml:"id"`
	Links        []AtomLink        `xml:"link"`
	Published    time.Time         `xml:"published"`
	Rights       string            `xml:"rights"`
	Source       string            `xml:"source"`
	Summary      atomText          `xml:"summary"`
	Title        string            `xml:"title"`
	Updated      time.Time         `xml:"updated"`

	// Dexter specific attributes
	feedID index.DexID
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

// SetSubscriptionID sets the given ID to the feed.
func (af *AtomFeed) SetSubscriptionID(id index.DexID) {
	af.subscriptionID = id
}

// ID implements the Feed interface.
func (af *AtomFeed) ID() string {
	return af.ID_
}

// Title implements the Feed interface.
func (af *AtomFeed) Title() string {
	return af.Title_
}

// SubscriptionID implements the Feed interface.
func (af *AtomFeed) SubscriptionID() index.DexID {
	return af.subscriptionID
}

// SetFeedID sets the given ID to the entry.
func (ae *AtomEntry) SetFeedID(id index.DexID) {
	ae.feedID = id
}
