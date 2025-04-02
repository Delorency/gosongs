package groupservice

import (
	"main/internal/models"
	"main/internal/schemes"
)

func (s *groupService) List(p *schemes.Pagination) (*[]models.Group, error) {
	return s.repo.List(p)
}

func (r *groupService) Retrieve(id uint) (*models.Group, error) {
	return r.repo.Retrieve(id)
}
