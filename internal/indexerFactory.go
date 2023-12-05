package internal

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/marcopeocchi/mayoi/internal/animetime"
	"github.com/marcopeocchi/mayoi/internal/domain"
	"github.com/marcopeocchi/mayoi/internal/nyaa"
	"github.com/marcopeocchi/mayoi/internal/registry"
)

func IndexerFactory(url string, db *sql.DB, r *registry.Registry, mux *http.ServeMux) (domain.Indexer, error) {
	switch url {
	case
		"https://nyaa.si/?page=rss",
		"https://nyaa.land/?page=rss",
		"https://nyaa.nocensor.cloud/?page=rss":
		indexer := nyaa.Module(db, r, mux, url)
		return indexer, nil
	case
		"https://animetime.cc/rss",
		"https://animetime.cc/rss/anime":
		indexer := animetime.Module(db, r, mux, url)
		return indexer, nil
	default:
		return nil, errors.New("no indexer implemented for this url")
	}
}
