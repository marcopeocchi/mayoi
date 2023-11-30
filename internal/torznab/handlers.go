package torznab

import (
	"encoding/xml"
	"net/http"
)

type TorznabHandler struct {
	repo *TorznabRepository
	mux  *http.ServeMux
}

func NewTorznabHandler(r *TorznabRepository, mux *http.ServeMux) *TorznabHandler {
	return &TorznabHandler{
		repo: r,
		mux:  mux,
	}
}

func (h *TorznabHandler) ApplyRoutes() {
	h.mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
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
