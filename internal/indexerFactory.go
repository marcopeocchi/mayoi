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

type IndexerFactoryParams struct {
	URL string
	DB  *sql.DB
	Reg *registry.Registry
	Mux *http.ServeMux
}

func IndexerFactory(args *IndexerFactoryParams) (domain.Indexer, error) {
	switch args.URL {
	case
		"https://nyaa.si/?page=rss",
		"https://nyaa.land/?page=rss",
		"https://nyaa.nocensor.cloud/?page=rss":
		indexer := nyaa.Module(args.DB, args.Reg, args.Mux, args.URL)
		return indexer, nil
	case
		"https://animetime.cc/rss",
		"https://animetime.cc/rss/anime":
		indexer := animetime.Module(args.DB, args.Reg, args.Mux, args.URL)
		return indexer, nil
	default:
		return nil, errors.New("no indexer implemented for this url")
	}
}
