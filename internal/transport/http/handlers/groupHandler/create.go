package grouphandler

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	models.Group
}

func (gh *groupHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request Request

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Print(err)
	}

	obj := models.Group{Name: request.Name}

	if err := gh.service.Create(&obj); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(obj)
}
