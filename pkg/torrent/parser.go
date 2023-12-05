package torrent

import (
	"crypto/sha1"
	"fmt"
	"io"
	"time"

	"github.com/zeebo/bencode"
)

type Torrent struct {
	CreatedAt time.Time
	InfoHash  string
	Length    int64
}

type Info struct {
	Name   string `bencode:"name"`
	Length int64  `bencode:"length"`
}

type Metadata struct {
	CreatedAt int64              `bencode:"creation date"`
	Info      bencode.RawMessage `bencode:"info"`
}

func Parse(r io.Reader) (*Torrent, error) {
	m := &Metadata{}

	if err := bencode.NewDecoder(r).Decode(m); err != nil {
		return nil, err
	}

	info := &Info{}
	if err := bencode.DecodeBytes(m.Info, info); err != nil {
		return nil, err
	}

	infoHash := sha1.New()
	infoHash.Write(m.Info)

	return &Torrent{
		CreatedAt: time.Unix(m.CreatedAt, 0),
		InfoHash:  fmt.Sprintf("%x", infoHash.Sum(nil)),
		Length:    info.Length,
	}, nil
}
