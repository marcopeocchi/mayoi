-- name: InsertFeed :exec
INSERT INTO animetime (
  guid, title, link, category, pubDate, infohash, created_at, size, seeders, peers
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetFeedsByTitle :many
SELECT * FROM animetime
WHERE title LIKE ? COLLATE NOCASE
ORDER BY created_at;

-- name: GetLatestFeeds :many
SELECT * FROM animetime
ORDER BY created_at
LIMIT 50;
