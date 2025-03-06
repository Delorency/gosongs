package grouprepo

import "main/internal/models"

func (r *GroupRepo) Create(group *models.Group) (uint, error) {
	err := r.db.Create(group).Error

	return group.ID, err
}
