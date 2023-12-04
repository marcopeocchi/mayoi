package nyaa

import (
	"context"
	"database/sql"
	"encoding/xml"
	"log/slog"
	"net/http"
	"time"

	"github.com/marcopeocchi/mayoi/internal/domain"
	"github.com/marcopeocchi/mayoi/internal/nyaa/db"
	"github.com/marcopeocchi/mayoi/pkg/useragent"
)

const indexerName = "Nyaa"

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

	for _, item := range feed.Channel.Items {
		if !isAnime(&item) {
			continue
		}
		q.InsertFeed(ctx, db.InsertFeedParams{
			Guid:     item.GUID,
			Title:    item.Title,
			Link:     item.Link,
			Category: "5070",
			Pubdate:  item.PubDate,
			Infohash: sql.NullString{
				String: item.InfoHash,
				Valid:  true,
			},
			Size: convertSize(item.Size),
			Seeders: sql.NullInt64{
				Int64: item.Seeders,
				Valid: true,
			},
			Peers: sql.NullInt64{
				Int64: item.Leechers,
				Valid: true,
			},
			CreatedAt: time.Now(),
		})
	}

	return nil
}

func (r *RSSIndexer) AutoIndex(ctx context.Context, d time.Duration) error {
	for {
		slog.Info(
			"Scanning",
			slog.String("indexer", indexerName),
			slog.String("url", r.url),
		)
		r.Index(ctx)
		time.Sleep(d)
	}
}
