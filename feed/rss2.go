package feed

import (
	"encoding/xml"
	"time"

	"github.com/toru/dexter/index"
)

// RSS2Time is a custom type that embeds the standard time.Time.
// Purpose of this type is to implement a custom XML node parser.
type RSS2Time struct {
	time.Time
}

// UnmarshalXML implements the xml.Unmarshaler interface.
// A generous implementation that allows multiple format strings.
func (rt *RSS2Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	var tmp string
	var parsed time.Time

	d.DecodeElement(&tmp, &start)
	fmts := []string{time.RFC1123Z, time.RFC1123, time.RFC822Z, time.RFC822}
	for _, f := range fmts {
		parsed, err = time.Parse(f, tmp)
		if err == nil {
			rt.Time = parsed
			break
		}
	}
	return err
}

type RSS2Channel struct {
	// Required Fields
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`

	// Optional Fields
	Language       string   `xml:"language"`
	Copyright      string   `xml:"copyright"`
	ManagingEditor string   `xml:"managingEditor"`
	WebMaster      string   `xml:"webMaster"`
	PubDate        RSS2Time `xml:"pubDate"`
	LastBuildDate  RSS2Time `xml:"lastBuildDate"`
	Category       []string `xml:"category"`
	Generator      string   `xml:"generator"`
	Docs           string   `xml:"docs"`
	Cloud          string   `xml:"cloud"`
	TTL            uint     `xml:"ttl"`
	Image          string   `xml:"image"`
	// TODO: TextInput
	// TODO: SkipHours
	// TODO: SkipDays
}

type RSS2Feed struct {
	XMLName xml.Name    `xml:"rss"`
	Version string      `xml:"version,attr"`
	Channel RSS2Channel `xml:"channel"`

	// Dexter specific attributes
	subscriptionID index.DexID
}

// SetSubscriptionID sets the given ID to the feed.
func (rf *RSS2Feed) SetSubscriptionID(id index.DexID) {
	rf.subscriptionID = id
}
