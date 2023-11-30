package management

import "net/http"

func Module(mux *http.ServeMux) {
	handler := NewManagementHandler(mux)

	handler.ApplyRoutes()
}
