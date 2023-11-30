package torznab

import (
	"database/sql"
	"net/http"
)

func Module(db *sql.DB, mux *http.ServeMux) {
	repo := ProvideRepository(db)
	handler := ProvideHandler(repo, mux)

	handler.ApplyRoutes()
}
