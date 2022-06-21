package server

import (
	"context"
	"ldap-integration/src/core"
	"ldap-integration/src/user"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, userService user.Service) http.Handler {
	r := mux.NewRouter()
	r.Use(BasicMiddleware)

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		core.ResponseOk(w, "API is running")
	}).Methods("GET")

	user.SetRoutes(ctx, r, userService)

	return r
}
