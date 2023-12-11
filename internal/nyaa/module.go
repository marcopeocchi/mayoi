package nyaa

import (
	"database/sql"
	"net/http"

	"github.com/marcopeocchi/mayoi/internal/domain"
	nyaaDB "github.com/marcopeocchi/mayoi/internal/nyaa/db"
	"github.com/marcopeocchi/mayoi/internal/registry"
)

func Module(db *sql.DB, r *registry.Registry, mux *http.ServeMux, url string) domain.Indexer {
	const path = "/nyaa/api"

	r.Set(url, path)

	var (
		repository = NewRepository(db)
		handler    = NewHandler(repository, mux, path)
		indexer    = NewIndexer(url, db)
	)

	nyaaDB.CreateTable(db)
	handler.ApplyRoutes()

	return indexer
}
