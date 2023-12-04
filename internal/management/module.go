package management

import (
	"net/http"

	"github.com/marcopeocchi/mayoi/internal/registry"
)

func Module(mux *http.ServeMux, reg *registry.Registry) {
	var (
		repository = NewReposityory()
		handler    = NewManagementHandler(mux, repository, reg)
	)
	handler.ApplyRoutes()
}
