package feed

import (
	"encoding/xml"
	"time"
)

type RSS2Channel struct {
	// Required Fields
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`

	// Optional Fields
	Language       string    `xml:"language"`
	Copyright      string    `xml:"copyright"`
	ManagingEditor string    `xml:"managingEditor"`
	WebMaster      string    `xml:"webMaster"`
	PubDate        time.Time `xml:"pubDate"`
	LastBuildDate  time.Time `xml:"lastBuildDate"`
	Category       []string  `xml:"category"`
	Generator      string    `xml:"generator"`
	Docs           string    `xml:"docs"`
	Cloud          string    `xml:"cloud"`
	TTL            uint      `xml:"ttl"`
	Image          string    `xml:"image"`
	// TODO: TextInput
	// TODO: SkipHours
	// TODO: SkipDays
}

type RSS2Feed struct {
	XMLName xml.Name    `xml:"rss"`
	Version string      `xml:"version,attr"`
	Channel RSS2Channel `xml:"channel"`
}
