package domain

import (
	"context"

	"github.com/marcopeocchi/mayoi/internal/rss"
)

type TorznabRepository interface {
	Caps() (*rss.ApiCapabilities, error)
	Search(ctx context.Context, query string) (*rss.Feed, error)
}
