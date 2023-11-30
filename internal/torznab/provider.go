package torznab

import (
	"database/sql"
	"net/http"
	"sync"
)

var (
	repo    *TorznabRepository
	handler *TorznabHandler

	repoOnce    sync.Once
	handlerOnce sync.Once
)

func ProvideRepository(db *sql.DB) *TorznabRepository {
	repoOnce.Do(func() {
		repo = NewTornzabRepository(db)
	})
	return repo
}

func ProvideHandler(repo *TorznabRepository, mux *http.ServeMux) *TorznabHandler {
	handlerOnce.Do(func() {
		handler = NewTorznabHandler(repo, mux)
	})
	return handler
}
