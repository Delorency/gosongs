package groupdb

import (
	"main/internal/models"
)

type Update struct {
	Name string `json:"name"`
}

func (r *groupDB) Update(schema, Update, group *models.Group) error {
	err := r.db.Model(&group).Updates(schema).Error

	return err
}
