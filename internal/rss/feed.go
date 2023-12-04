package rss

import (
	"encoding/xml"
)

type Feed struct {
	XMLName      xml.Name `xml:"rss"`
	Version      string   `xml:"version,attr"`
	XMLNSAtom    string   `xml:"xmlns:atom,attr"`
	XMLNSTorznab string   `xml:"xmlns:torznab,attr"`
	Channel      Channel  `xml:"channel"`
	Items        []Item   `xml:"item"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	AtomLink    AtomLink
	Items       []Item `xml:"item"`
}

type AtomLink struct {
	XMLName xml.Name `xml:"atom:link"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

type Item struct {
	GUID         string `xml:"guid"`
	Title        string `xml:"title"`
	Link         string `xml:"link"`
	Category     string `xml:"category"`
	Size         int64  `xml:"size"`
	PubDate      string `xml:"pubDate"`
	Enclosure    TorznabEnclosure
	TorznabAttrs []TorznabAttr
}
