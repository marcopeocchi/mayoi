package management

import (
	"encoding/json"
	"net/http"

	"github.com/marcopeocchi/mayoi/internal/registry"
)

type ManagementHandler struct {
	mux        *http.ServeMux
	repository *Repository
	reg        *registry.Registry
}

func NewManagementHandler(mux *http.ServeMux, r *Repository, reg *registry.Registry) *ManagementHandler {
	return &ManagementHandler{
		mux:        mux,
		reg:        reg,
		repository: r,
	}
}

func (h *ManagementHandler) ApplyRoutes() {
	h.mux.HandleFunc("/management/db", func(w http.ResponseWriter, r *http.Request) {
		dbsize, err := h.repository.GetDatabaseSize()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(dbsize); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	h.mux.HandleFunc("/management/indexers", func(w http.ResponseWriter, r *http.Request) {
		if err := h.reg.JsonEncoder(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
