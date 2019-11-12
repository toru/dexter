package feed

import (
	"encoding/xml"
	"time"

	"github.com/toru/dexter/index"
)

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
	Entries_     []AtomEntry       `xml:"entry"`

	// Dexter specific attributes
	subscriptionID index.ID
}

// AtomEntry represents the "atom:entry" container element defined in
// RFC 4287 Section 4.1.2
type AtomEntry struct {
	Authors      []AtomAuthor      `xml:"author"`
	Categories   []AtomCategory    `xml:"category"`
	Content      atomText          `xml:"content"`
	Contributors []AtomContributor `xml:"contributor"`
	ID_          string            `xml:"id"`
	Links        []AtomLink        `xml:"link"`
	Published    time.Time         `xml:"published"`
	Rights       string            `xml:"rights"`
	Source       string            `xml:"source"`
	Summary_     atomText          `xml:"summary"`
	Title        string            `xml:"title"`
	Updated      time.Time         `xml:"updated"`

	// Dexter specific attributes
	feedID index.ID
}

// NewAtomFeed returns a new AtomFeed.
func NewAtomFeed() Feed {
	feed := &AtomFeed{}
	// TODO: Check if this instantiation is needed. Makes more
	// sense to set the pointer from the caller via SetSubID().
	feed.subscriptionID = index.GetIssuer().CreateID()
	return feed
}

// ParseAtomFeed parses the given byte slice as an AtomFeed.
func ParseAtomFeed(doc []byte) (Feed, error) {
	f := NewAtomFeed()
	if err := xml.Unmarshal(doc, f); err != nil {
		return nil, err
	}
	return f, nil
}

// SetSubscriptionID sets the given ID to the feed.
func (af *AtomFeed) SetSubscriptionID(id index.ID) {
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

// Format implements the Feed interface.
func (af *AtomFeed) Format() uint {
	return AtomFeedFormat
}

// SubscriptionID implements the Feed interface.
func (af *AtomFeed) SubscriptionID() index.ID {
	return af.subscriptionID
}

// Entries implements the Feed interface.
func (af *AtomFeed) Entries() []Entry {
	rv := make([]Entry, len(af.Entries_))
	for i := range af.Entries_ {
		rv[i] = &af.Entries_[i]
	}
	return rv
}

// SetFeedID implements the Entry interface.
func (ae *AtomEntry) SetFeedID(id index.ID) {
	ae.feedID = id
}

// FeedID implements the Entry interface.
func (ae *AtomEntry) FeedID() index.ID {
	return ae.feedID
}

// ID implements the Entry interface.
func (ae *AtomEntry) ID() string {
	return ae.ID_
}

// Summary implements the Entry interface.
func (ae *AtomEntry) Summary() string {
	return ae.Summary_.Value
}
