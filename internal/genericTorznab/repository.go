package generictorznab

import (
	"context"
	"database/sql"

	"github.com/marcopeocchi/mayoi/internal/rss"
)

type TorznabRepository struct {
	db *sql.DB
}

func NewTornzabRepository(db *sql.DB) *TorznabRepository {
	return &TorznabRepository{
		db: db,
	}
}

func (r *TorznabRepository) Caps() (*rss.ApiCapabilities, error) {
	caps := &rss.ApiCapabilities{}

	caps.Server.Version = "1.0"
	caps.Server.Title = "Mayoi"
	caps.Server.Strapline = "Mayoi Anime Indexer"

	caps.Retention.Days = "60"

	caps.Limits.Default = "100"
	caps.Limits.Max = "100"

	// Searching ------------------------------------------------------------- //

	caps.Searching.Search.Available = "yes"
	caps.Searching.Search.SupportedParams = "q"

	caps.Searching.TvSearch.Available = "yes"
	caps.Searching.TvSearch.SupportedParams = "q,season,ep"

	caps.Searching.AudioSearch.Available = "no"
	caps.Searching.AudioSearch.SupportedParams = "q"

	caps.Searching.MovieSearch.Available = "no"
	caps.Searching.MovieSearch.SupportedParams = "q"

	caps.Searching.BookSearch.Available = "no"
	caps.Searching.BookSearch.SupportedParams = "q"

	// Categories ------------------------------------------------------------ //

	caps.Categories.Category = []rss.TorznabCategory{
		{
			ID:   "5000",
			Name: "TV",
			Subcat: []rss.TorznabSubCategory{
				{
					ID:   "5070",
					Name: "TV/Anime",
				},
			},
		},
	}

	return caps, nil
}

func (r *TorznabRepository) Search(ctx context.Context, query string) (*rss.Feed, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	statement := `
		SELECT rowid, guid, title, link, category, pubDate, infohash 
		FROM feeds 
		WHERE title LIKE ? 
		COLLATE NOCASE 
		ORDER BY rowid DESC`

	if query == "" {
		statement = `
			SELECT rowid, guid, title, link, category, pubDate, infohash 
			FROM feeds 
			ORDER BY rowid DESC 
			LIMIT 100`
	}

	rows, err := r.db.QueryContext(ctx, statement, "%"+query+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feed rss.Feed
	feed.Version = "1.0"

	feed.Channel.Title = "Mayoi"
	feed.Channel.Description = "Mayoi Indexer"
	feed.Channel.AtomLink.Rel = "self"
	feed.Channel.AtomLink.Type = "application/rss+xml"

	feed.XMLNSAtom = "http://www.w3.org/2005/Atom"
	feed.XMLNSTorznab = "http://torznab.com/schemas/2015/feed"

	var items []rss.Item

	for rows.Next() {
		var id int64
		var item rss.Item
		var infohash string

		rows.Scan(
			&id,
			&item.GUID,
			&item.Title,
			&item.Link,
			&item.Category,
			&item.PubDate,
			&infohash,
		)

		item.Size = 377697088

		item.TorznabAttrs = []rss.TorznabAttr{
			{
				Name:  "category",
				Value: "5070",
			},
			{
				Name:  "tag",
				Value: "freeleech",
			},
			{
				Name:  "seeders",
				Value: "1",
			},
			{
				Name:  "peers",
				Value: "2",
			},
			{
				Name:  "infohash",
				Value: infohash,
			},
			{
				Name:  "downloadvolumefactor",
				Value: "0",
			},
			{
				Name:  "uploadvolumefactor",
				Value: "1",
			},
		}

		items = append(items, item)
	}

	feed.Channel.Items = items

	return &feed, nil
}
