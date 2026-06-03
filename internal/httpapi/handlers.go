package httpapi

import (
	"net/http"
)

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	response := map[string]string{
		"status": "ok",
	}

	writeJSON(w, http.StatusOK, response)
}
