package utils

import (
	"context"
	"database/sql"
)

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
