package groupdb

import (
	"main/internal/models"

	"gorm.io/gorm"
)

func (r *groupDB) Update(id uint, data *models.Group) (*models.Group, error) {
	obj := models.Group{Model: gorm.Model{ID: id}}

	if err := r.db.Model(&obj).Updates(data).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(&obj).Error; err != nil {
		return nil, err
	}

	return &obj, nil
}
