package db

import (
	"context"
	"database/sql"
)

func CreateTable(db *sql.DB) error {
	ctx := context.Background()

	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS nyaa (
		guid VARCHAR(255) UNIQUE NOT NULL,
		title VARCHAR(255) NOT NULL, 
		link TEXT NOT NULL, 
		category VARCHAR(4) NOT NULL, 
		pubdate VARCHAR(255) NOT NULL,
		infohash VARCHAR(255),
		created_at DATETIME NOT NULL,
		size INT NOT NULL,
		seeders INT,
		peers INT
	)`)

	return err
}
