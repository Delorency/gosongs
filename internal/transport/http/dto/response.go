package dto

import (
	"encoding/json"
	"net/http"
)

func NewResponse(s any, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(s)
	return
}
