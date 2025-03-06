package grouprepo

import (
	"main/internal/models"
	gs "main/internal/schemes/groupSchemes"
)

func (r *GroupRepo) Update(id uint, schema gs.Update, group *models.Group) error {
	err := r.db.Model(&group).Updates(schema).Error

	return err
}
