package torrent_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/marcopeocchi/mayoi/pkg/torrent"
)

func TestParsing(t *testing.T) {
	var (
		client = http.DefaultClient
		url    = "https://archlinux.org/releng/releases/2023.12.01/torrent/"
	)

	res, err := client.Get(url)
	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	tor, err := torrent.Parse(res.Body)
	if err != nil {
		t.Error(err)
	}

	if tor.InfoHash == "" {
		t.Error(errors.New("no infohash parsed"))
	}
	if tor.Length == 0 {
		t.Error(errors.New("cannot find file size"))
	}
}
