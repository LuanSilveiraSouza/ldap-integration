package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer() http.Handler {
	r := mux.NewRouter()
	r.Use(BasicMiddleware)

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("API is running")
	}).Methods("GET")

	return r
}
