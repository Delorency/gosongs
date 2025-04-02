package groupdb

import "main/internal/models"

func (r *groupDB) Create(data *models.Group) error {
	return r.db.Create(data).Error
}
