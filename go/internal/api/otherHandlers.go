package api

import (
	"encoding/json"
	"net/http"
)

func (s *server) handleTest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Ok")
	}
}
