package animetime

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/marcopeocchi/mayoi/internal/animetime/db"
	"github.com/marcopeocchi/mayoi/internal/domain"
	"github.com/marcopeocchi/mayoi/pkg/torrent"
	"github.com/marcopeocchi/mayoi/pkg/useragent"
)

const (
	baseUrl     = "https://animetime.cc"
	indexerName = "AnimeTime"
)

type RSSIndexer struct {
	url        string
	db         *sql.DB
	httpClient *http.Client
}

func NewIndexer(url string, db *sql.DB) domain.Indexer {
	return &RSSIndexer{
		url:        url,
		db:         db,
		httpClient: http.DefaultClient,
	}
}

func (r *RSSIndexer) scan() (*Feed, error) {
	req, err := http.NewRequest(http.MethodGet, r.url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", useragent.Default)

	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var feed Feed

	if err := xml.NewDecoder(res.Body).Decode(&feed); err != nil {
		return nil, err
	}

	return &feed, err
}

func (r *RSSIndexer) downloadTorrent(url string) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", useragent.Default)

	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func (r *RSSIndexer) Index(ctx context.Context) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	q := db.New(conn)

	feed, err := r.scan()
	if err != nil {
		return err
	}

	for _, item := range feed.Items {
		if !isAnime(&item) {
			continue
		}

		url := fmt.Sprintf("%s/download/%s", baseUrl, item.GUID)

		res, err := r.downloadTorrent(url)
		if err != nil {
			slog.Error(
				"cannot get torrent file",
				slog.String("url", url),
				slog.String("err", err.Error()),
			)
			continue
		}

		tor, err := torrent.Parse(res)
		if err != nil {
			slog.Error(
				"cannot parse torrent file",
				slog.String("url", url),
				slog.String("err", err.Error()),
			)
			continue
		}

		res.Close()

		q.InsertFeed(ctx, db.InsertFeedParams{
			Guid:     item.GUID,
			Title:    item.Title,
			Link:     url,
			Category: "5070",
			Pubdate:  item.PubDate,
			Infohash: sql.NullString{
				String: tor.InfoHash,
				Valid:  true,
			},
			Size: tor.Length,
			Seeders: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			Peers: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			CreatedAt: time.Now(),
		})
	}

	time.Sleep(time.Second)

	return nil
}

func (r *RSSIndexer) AutoIndex(ctx context.Context, d time.Duration) error {
	for {
		slog.Info(
			"Scanning",
			slog.String("indexer", indexerName),
			slog.String("url", r.url),
		)
		if err := r.Index(ctx); err != nil {
			slog.Error(
				"Failed scanning",
				slog.String("indexer", indexerName),
				slog.String("url", r.url),
				slog.String("err", err.Error()),
			)
			return err
		}
		time.Sleep(d)
	}
}
