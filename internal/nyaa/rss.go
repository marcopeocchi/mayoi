package nyaa

import (
	"encoding/xml"
)

type Feed struct {
	XMLName   xml.Name `xml:"rss"`
	Version   string   `xml:"version,attr"`
	XMLNSAtom string   `xml:"xmlns:atom,attr"`
	XMLNSNyaa string   `xml:"xmlns:nyaa,attr"`
	Channel   Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	GUID       string `xml:"guid"`
	Title      string `xml:"title"`
	Link       string `xml:"link"`
	PubDate    string `xml:"pubDate"`
	Seeders    int64  `xml:"seeders"`
	Leechers   int64  `xml:"leechers"`
	Downloads  string `xml:"downloads"`
	InfoHash   string `xml:"infoHash"`
	CategoryId string `xml:"categoryId"`
	Category   string `xml:"category"`
	Size       string `xml:"size"`
	Comments   string `xml:"comments"`
	Trusted    string `xml:"trusted"`
	Remake     string `xml:"remake"`
}
