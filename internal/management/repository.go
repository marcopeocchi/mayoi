package management

import (
	"context"
	"database/sql"
	"os"
	"strings"

	"github.com/marcopeocchi/mayoi/internal/registry"
	"github.com/marcopeocchi/mayoi/pkg/config"
)

type Repository struct {
	reg *registry.Registry
	db  *sql.DB
}

func NewReposityory(db *sql.DB, reg *registry.Registry) *Repository {
	return &Repository{
		db:  db,
		reg: reg,
	}
}

func (r *Repository) GetDatabaseSize() (int64, error) {
	stat, err := os.Stat(config.Instance().Database)
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}

func (r *Repository) GetIndexedEntries(ctx context.Context) (int64, error) {
	var (
		total    int64
		indexers = r.reg.Keys()
	)

	conn, err := r.db.Conn(ctx)
	if err != nil {
		return 0, err
	}

	for _, name := range *indexers {
		var (
			count     int64
			tableName = strings.TrimPrefix(strings.TrimSuffix(name, "/api"), "/")
		)

		row := conn.QueryRowContext(ctx, "SELECT COUNT(1) FROM "+tableName)

		err := row.Scan(&count)
		if err != nil {
			return 0, err
		}

		total += count
	}

	return total, nil
}
