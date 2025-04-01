package groupdb

import "main/internal/models"

func (r *groupDB) Create(group *models.Group) error {
	return r.db.Create(group).Error
}
