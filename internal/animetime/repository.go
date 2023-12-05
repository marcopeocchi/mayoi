package animetime

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/marcopeocchi/mayoi/internal/animetime/db"
	"github.com/marcopeocchi/mayoi/internal/rss"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Caps() (*rss.ApiCapabilities, error) {
	caps := &rss.ApiCapabilities{}

	caps.Server.Version = "1.0"
	caps.Server.Title = "Mayoi (AnimeTime)"
	caps.Server.Strapline = "Mayoi Indexer (AnimeTime)"

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

func (r *Repository) Search(ctx context.Context, query string) (*rss.Feed, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	q := db.New(conn)

	var feeds []db.Animetime

	if query != "" {
		feeds, err = q.GetFeedsByTitle(ctx, "%"+query+"%")
		if err != nil {
			return nil, err
		}
	} else {
		feeds, err = q.GetLatestFeeds(ctx)
		if err != nil {
			return nil, err
		}
	}

	var (
		feed  rss.Feed
		items = make([]rss.Item, len(feeds))
	)

	feed.Version = "1.0"

	feed.Channel.Title = "Mayoi"
	feed.Channel.Description = "Mayoi Indexer"
	feed.Channel.AtomLink.Rel = "self"
	feed.Channel.AtomLink.Type = "application/rss+xml"

	feed.XMLNSAtom = "http://www.w3.org/2005/Atom"
	feed.XMLNSTorznab = "http://torznab.com/schemas/2015/feed"

	for i, feed := range feeds {
		items[i].GUID = feed.Guid
		items[i].Category = feed.Category
		items[i].Link = feed.Link
		items[i].PubDate = feed.Pubdate
		items[i].Size = feed.Size
		items[i].Title = feed.Title

		items[i].Enclosure.Url = feed.Link
		items[i].Enclosure.Length = feed.Size
		items[i].Enclosure.Type = "application/x-bittorrent"

		items[i].TorznabAttrs = []rss.TorznabAttr{
			{
				Name:  "category",
				Value: feed.Category,
			},
			{
				Name:  "tag",
				Value: "freeleech",
			},
			{
				Name:  "seeders",
				Value: strconv.FormatInt(feed.Peers.Int64, 10),
			},
			{
				Name:  "peers",
				Value: strconv.FormatInt(feed.Peers.Int64, 10),
			},
			{
				Name:  "infohash",
				Value: feed.Infohash.String,
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
	}

	feed.Channel.Items = items

	return &feed, nil
}
