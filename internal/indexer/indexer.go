package indexer

import (
	"context"
	"database/sql"
	"encoding/xml"
	"log"
	"net/http"
	"time"

	"github.com/marcopeocchi/mayoi/internal/rss"
	"github.com/marcopeocchi/mayoi/pkg/useragent"
	"github.com/marcopeocchi/mayoi/pkg/utils"
)

type RSSIndexer struct {
	url        string
	db         *sql.DB
	httpClient *http.Client
}

func New(url string, db *sql.DB) *RSSIndexer {
	return &RSSIndexer{
		url:        url,
		db:         db,
		httpClient: http.DefaultClient,
	}
}

func (r *RSSIndexer) Scan() (*rss.Feed, error) {
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

	var feed rss.Feed

	if err := xml.NewDecoder(res.Body).Decode(&feed); err != nil {
		return nil, err
	}

	if len(feed.Channel.Items) == 0 {
		feed.Channel.Items = feed.Items
		feed.Items = nil
	}

	for i := range feed.Channel.Items {
		feed.Channel.Items[i].Category = "5070"
	}

	return &feed, err
}

func (r *RSSIndexer) Index(ctx context.Context) error {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	feed, err := r.Scan()
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Items {
		conn.ExecContext(
			ctx,
			`INSERT INTO feeds 
			(guid, title, link, category, pubDate, infohash, createdAt) 
			VALUES 
			(?, ?, ?, ?, ?, ?, ?)`,
			item.GUID,
			item.Title,
			item.Link,
			item.Category,
			item.PubDate,
			utils.InfoHashFromMagnet(item.Link),
			time.Now(),
		)
	}

	return nil
}

func (r *RSSIndexer) AutoIndex(d time.Duration) error {
	for {
		log.Printf("Scanning %s\n", r.url)
		r.Index(context.Background())
		time.Sleep(d)
	}
}
