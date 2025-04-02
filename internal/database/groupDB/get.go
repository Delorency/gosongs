package groupdb

import (
	"main/internal/models"
	"main/internal/schemes"

	"gorm.io/gorm"
)

func (r *groupDB) List(p *schemes.Pagination) (*[]models.Group, error) {
	var groups []models.Group

	err := r.db.Where(models.Group{}).Limit(p.Limit).Offset((p.Page - 1) * p.Limit).Order("created_at desc").Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return &groups, err
}

func (r *groupDB) Retrieve(id uint) (*models.Group, error) {
	obj := models.Group{Model: gorm.Model{ID: id}}

	err := r.db.First(&obj).Error

	if err != nil {
		return nil, err
	}

	return &obj, err
}
