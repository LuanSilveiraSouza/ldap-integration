package login

import (
	"context"
	"encoding/json"
	"ldap-integration/src/core"
	"net/http"

	"golang.org/x/oauth2"
)

type Service interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type service struct {
	repository Repository
	dex        oauth2.Config
}

func NewService(rep Repository, dex oauth2.Config) Service {
	return &service{repository: rep, dex: dex}
}

func (s service) Login(w http.ResponseWriter, r *http.Request) {
	var l LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		core.ResponseBadRequest(w, "Invalid body")
	}

	user, err := s.repository.Authenticate(l.Email)
	if err != nil {
		core.ResponseBadRequest(w, "Invalid Login")
	}

	if _, err := s.dex.PasswordCredentialsToken(context.Background(), l.Email, l.Password); err != nil {
		core.ResponseBadRequest(w, "Invalid Login")
	} else {

	}
}
