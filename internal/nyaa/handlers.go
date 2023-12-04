package nyaa

import (
	"encoding/xml"
	"net/http"
)

type NyaaHandler struct {
	repo *Repository
	mux  *http.ServeMux
	path string
}

func NewNyaaHandler(r *Repository, mux *http.ServeMux, path string) *NyaaHandler {
	return &NyaaHandler{
		repo: r,
		mux:  mux,
		path: path,
	}
}

func (h *NyaaHandler) ApplyRoutes() {
	h.mux.HandleFunc(h.path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")

		fn := r.URL.Query().Get("t")
		query := r.URL.Query().Get("q")

		switch fn {
		case "search", "tvsearch", "":
			feed, err := h.repo.Search(r.Context(), query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := xml.NewEncoder(w).Encode(feed); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "caps":
			caps, err := h.repo.Caps()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := xml.NewEncoder(w).Encode(caps); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})
}
