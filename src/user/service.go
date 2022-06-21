package user

import (
	"ldap-integration/src/core"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service interface {
	ListUserTypes(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type service struct {
	repository Repository
}

func NewService(rep Repository) Service {
	return &service{repository: rep}
}

func (s service) ListUserTypes(w http.ResponseWriter, r *http.Request) {
	userTypes, err := s.repository.ListUserTypes()

	if err != nil {
		core.ResponseInternalError(w, err.Error())
	}

	core.ResponseOk(w, userTypes)
}

func (s service) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		core.ResponseBadRequest(w, err.Error())
		return
	}

	user, err := s.repository.Get(int64(id))
	if err != nil {
		core.ResponseInternalError(w, err.Error())
	}

	core.ResponseOk(w, user)
}

func (s service) List(w http.ResponseWriter, r *http.Request) {
	users, err := s.repository.List()

	if err != nil {
		core.ResponseInternalError(w, err.Error())
	}

	core.ResponseOk(w, users)
}
