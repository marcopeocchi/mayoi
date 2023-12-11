-- name: InsertFeed :exec
INSERT INTO nyaa (
  guid, title, link, category, pubDate, infohash, created_at, size, seeders, peers
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetFeedsByTitle :many
SELECT * FROM nyaa
WHERE title LIKE ? COLLATE NOCASE
ORDER BY created_at DESC;

-- name: GetLatestFeeds :many
SELECT * FROM nyaa
ORDER BY created_at DESC
LIMIT 50;
