package nyaa

import (
	"database/sql"
	"net/http"

	"github.com/marcopeocchi/mayoi/internal/domain"
	nyaaDB "github.com/marcopeocchi/mayoi/internal/nyaa/db"
)

func Module(db *sql.DB, mux *http.ServeMux, url string) domain.Indexer {
	var (
		repository = NewRepository(db)
		handler    = NewNyaaHandler(repository, mux)
		indexer    = NewIndexer(url, db)
	)

	nyaaDB.AutoMigrate(db)
	handler.ApplyRoutes()

	return indexer
}
