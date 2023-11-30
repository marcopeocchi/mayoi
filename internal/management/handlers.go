package management

import (
	"encoding/json"
	"net/http"

	"github.com/marcopeocchi/mayoi/pkg/config"
	"github.com/marcopeocchi/mayoi/pkg/utils"
)

type ManagementHandler struct {
	mux *http.ServeMux
}

func NewManagementHandler(mux *http.ServeMux) *ManagementHandler {
	return &ManagementHandler{
		mux: mux,
	}
}

func (h *ManagementHandler) ApplyRoutes() {
	h.mux.HandleFunc("/management/db", func(w http.ResponseWriter, r *http.Request) {
		dbsize, err := utils.GetDatabaseSize()
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
		if err := json.NewEncoder(w).Encode(config.Instance().Indexers); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
