package feed

import (
	"encoding/xml"
)

type RSS2Feed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
}
