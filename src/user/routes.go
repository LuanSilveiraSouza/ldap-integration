package user

import (
	"context"

	"github.com/gorilla/mux"
)

func SetRoutes(ctx context.Context, r *mux.Router, userService Service) {
	r.HandleFunc("/api/users/types", userService.ListUserTypes).Methods("GET")
	r.HandleFunc("/api/users/{id}", userService.Get).Methods("GET")
	r.HandleFunc("/api/users", userService.List).Methods("GET")
}
