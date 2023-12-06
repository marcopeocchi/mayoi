// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getFeedsByTitle = `-- name: GetFeedsByTitle :many
SELECT guid, title, link, category, pubdate, infohash, created_at, size, seeders, peers FROM animetime
WHERE title LIKE ? COLLATE NOCASE
ORDER BY created_at
`

func (q *Queries) GetFeedsByTitle(ctx context.Context, title string) ([]Animetime, error) {
	rows, err := q.db.QueryContext(ctx, getFeedsByTitle, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Animetime
	for rows.Next() {
		var i Animetime
		if err := rows.Scan(
			&i.Guid,
			&i.Title,
			&i.Link,
			&i.Category,
			&i.Pubdate,
			&i.Infohash,
			&i.CreatedAt,
			&i.Size,
			&i.Seeders,
			&i.Peers,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLatestFeeds = `-- name: GetLatestFeeds :many
SELECT guid, title, link, category, pubdate, infohash, created_at, size, seeders, peers FROM animetime
ORDER BY created_at
LIMIT 50
`

func (q *Queries) GetLatestFeeds(ctx context.Context) ([]Animetime, error) {
	rows, err := q.db.QueryContext(ctx, getLatestFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Animetime
	for rows.Next() {
		var i Animetime
		if err := rows.Scan(
			&i.Guid,
			&i.Title,
			&i.Link,
			&i.Category,
			&i.Pubdate,
			&i.Infohash,
			&i.CreatedAt,
			&i.Size,
			&i.Seeders,
			&i.Peers,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertFeed = `-- name: InsertFeed :exec
INSERT INTO animetime (
  guid, title, link, category, pubDate, infohash, created_at, size, seeders, peers
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type InsertFeedParams struct {
	Guid      string
	Title     string
	Link      string
	Category  string
	Pubdate   string
	Infohash  sql.NullString
	CreatedAt time.Time
	Size      int64
	Seeders   sql.NullInt64
	Peers     sql.NullInt64
}

func (q *Queries) InsertFeed(ctx context.Context, arg InsertFeedParams) error {
	_, err := q.db.ExecContext(ctx, insertFeed,
		arg.Guid,
		arg.Title,
		arg.Link,
		arg.Category,
		arg.Pubdate,
		arg.Infohash,
		arg.CreatedAt,
		arg.Size,
		arg.Seeders,
		arg.Peers,
	)
	return err
}