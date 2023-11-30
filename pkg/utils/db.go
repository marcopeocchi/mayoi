package utils

import (
	"context"
	"database/sql"
)

func AutoMigrate(ctx context.Context, db *sql.DB) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS feeds(
		guid VARCHAR(255) UNIQUE NOT NULL,
		title VARCHAR(255) NOT NULL, 
		link VARCHAR(255) NOT NULL, 
		category CHAR(4) NOT NULL, 
		pubDate VARCHAR(255),
		infohash VARCHAR(255),
		createdAt DATE NOT NULL
	)`)

	return err
}

func PruneDatabase(ctx context.Context, db *sql.DB) error {
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.ExecContext(
		ctx,
		"DELETE FROM feeds WHERE createdAt <= date('now','-30 day')",
	)

	return err
}
