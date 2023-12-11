package management

import (
	"database/sql"
	"net/http"

	"github.com/marcopeocchi/mayoi/internal/registry"
)

func Module(mux *http.ServeMux, db *sql.DB, reg *registry.Registry) {
	var (
		repository = NewReposityory(db, reg)
		handler    = NewManagementHandler(mux, repository, reg)
	)
	handler.ApplyRoutes()
}
