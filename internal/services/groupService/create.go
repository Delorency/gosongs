package groupservice

import (
	"main/internal/models"
)

type Create struct {
	Name string `json:"name"`
}

func (s *groupService) Create(obj *models.Group) error {
	return s.repo.Create(obj)
}
