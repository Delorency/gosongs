package grouphandler

import (
	groupservice "main/internal/services/groupService"
	"net/http"
)

type GroupHandlerI interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type groupHandler struct {
	service groupservice.GroupServiceI
}

func NewGroupHandler(service groupservice.GroupServiceI) GroupHandlerI {
	return &groupHandler{service}
}
