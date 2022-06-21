package core

import (
	"encoding/json"
	"net/http"
)

func ResponseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

func ResponseBadRequest(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(body)
}

func ResponseInternalError(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(body)
}
