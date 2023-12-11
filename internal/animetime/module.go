package animetime

import (
	"database/sql"
	"net/http"

	animetimeDB "github.com/marcopeocchi/mayoi/internal/animetime/db"
	"github.com/marcopeocchi/mayoi/internal/domain"
	"github.com/marcopeocchi/mayoi/internal/registry"
)

func Module(db *sql.DB, r *registry.Registry, mux *http.ServeMux, url string) domain.Indexer {
	const path = "/animetime/api"

	r.Set(url, path)

	var (
		repository = NewRepository(db)
		handler    = NewHandler(repository, mux, path)
		indexer    = NewIndexer(url, db)
	)

	animetimeDB.CreateTable(db)
	handler.ApplyRoutes()

	return indexer
}
