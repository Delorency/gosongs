package groupservice

import "main/internal/models"

func (s *groupService) Update(id uint, data *models.Group) (*models.Group, error) {
	return s.repo.Update(id, data)
}
