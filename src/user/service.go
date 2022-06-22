package user

import (
	"ldap-integration/src/core"
	"ldap-integration/src/ldap"
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
	repository  Repository
	ldapService ldap.Service
}

func NewService(rep Repository, ldapSrv ldap.Service) Service {
	return &service{repository: rep, ldapService: ldapSrv}
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

	entry, err := s.ldapService.Get("admin@example.com")
	if err != nil {
		core.ResponseInternalError(w, err.Error())
	}

	entry.Print()

	core.ResponseOk(w, user)
}

func (s service) List(w http.ResponseWriter, r *http.Request) {
	users, err := s.repository.List()

	if err != nil {
		core.ResponseInternalError(w, err.Error())
	}

	core.ResponseOk(w, users)
}
