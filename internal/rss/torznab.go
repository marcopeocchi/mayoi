package rss

import "encoding/xml"

type TorznabAttr struct {
	XMLName xml.Name `xml:"torznab:attr"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type ApiCapabilities struct {
	XMLName xml.Name `xml:"caps"`
	Text    string   `xml:",chardata"`
	Server  struct {
		Text      string `xml:",chardata"`
		Version   string `xml:"version,attr"`
		Title     string `xml:"title,attr"`
		Strapline string `xml:"strapline,attr"`
		Email     string `xml:"email,attr"`
		URL       string `xml:"url,attr"`
		Image     string `xml:"image,attr"`
	} `xml:"server"`
	Limits struct {
		Text    string `xml:",chardata"`
		Max     string `xml:"max,attr"`
		Default string `xml:"default,attr"`
	} `xml:"limits"`
	Retention struct {
		Text string `xml:",chardata"`
		Days string `xml:"days,attr"`
	} `xml:"retention"`
	Registration struct {
		Text      string `xml:",chardata"`
		Available string `xml:"available,attr"`
		Open      string `xml:"open,attr"`
	} `xml:"registration"`
	Searching struct {
		Text   string `xml:",chardata"`
		Search struct {
			Text            string `xml:",chardata"`
			Available       string `xml:"available,attr"`
			SupportedParams string `xml:"supportedParams,attr"`
		} `xml:"search"`
		TvSearch struct {
			Text            string `xml:",chardata"`
			Available       string `xml:"available,attr"`
			SupportedParams string `xml:"supportedParams,attr"`
		} `xml:"tv-search"`
		MovieSearch struct {
			Text            string `xml:",chardata"`
			Available       string `xml:"available,attr"`
			SupportedParams string `xml:"supportedParams,attr"`
		} `xml:"movie-search"`
		AudioSearch struct {
			Text            string `xml:",chardata"`
			Available       string `xml:"available,attr"`
			SupportedParams string `xml:"supportedParams,attr"`
		} `xml:"audio-search"`
		BookSearch struct {
			Text            string `xml:",chardata"`
			Available       string `xml:"available,attr"`
			SupportedParams string `xml:"supportedParams,attr"`
		} `xml:"book-search"`
	} `xml:"searching"`
	Categories struct {
		Text     string            `xml:",chardata"`
		Category []TorznabCategory `xml:"category"`
	} `xml:"categories"`
	Groups struct {
		Text  string `xml:",chardata"`
		Group struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr"`
			Name        string `xml:"name,attr"`
			Description string `xml:"description,attr"`
			Lastupdate  string `xml:"lastupdate,attr"`
		} `xml:"group"`
	} `xml:"groups"`
	Genres struct {
		Text  string `xml:",chardata"`
		Genre struct {
			Text       string `xml:",chardata"`
			ID         string `xml:"id,attr"`
			Categoryid string `xml:"categoryid,attr"`
			Name       string `xml:"name,attr"`
		} `xml:"genre"`
	} `xml:"genres"`
	Tags struct {
		Text string `xml:",chardata"`
		Tag  []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			Description string `xml:"description,attr"`
		} `xml:"tag"`
	} `xml:"tags"`
}

type TorznabCategory struct {
	Text   string               `xml:",chardata"`
	ID     string               `xml:"id,attr"`
	Name   string               `xml:"name,attr"`
	Subcat []TorznabSubCategory `xml:"subcat"`
}

type TorznabSubCategory struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type TorznabEnclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  int64    `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}
