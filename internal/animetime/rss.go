package animetime

import (
	"encoding/xml"
)

type Feed struct {
	XMLName   xml.Name `xml:"rss"`
	Version   string   `xml:"version,attr"`
	XMLNSAtom string   `xml:"xmlns:atom,attr"`
	Channel   Channel  `xml:"channel"`
	Items     []Item   `xml:"item"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

type Item struct {
	GUID     string `xml:"guid"`
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	PubDate  string `xml:"pubDate"`
	Category string `xml:"category"`
}
