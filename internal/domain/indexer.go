package domain

import (
	"context"
	"time"
)

type Indexer interface {
	Index(ctx context.Context) error
	AutoIndex(ctx context.Context, d time.Duration) error
}
