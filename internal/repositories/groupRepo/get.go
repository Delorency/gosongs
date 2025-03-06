package grouprepo

import (
	"main/internal/models"
	"main/internal/schemes"
)

func (r *GroupRepo) Get(group *models.Group, p *schemes.Pagination) ([]models.Group, error) {
	var groups []models.Group

	err := r.db.Where(group).Limit(p.Limit).Offset(p.Page).Find(groups).Error

	if err != nil {
		return nil, err
	}

	return groups, err
}

func (r *GroupRepo) Retrieve(group *models.Group) (models.Group, error) {
	err := r.db.First(group).Error

	return *group, err
}
