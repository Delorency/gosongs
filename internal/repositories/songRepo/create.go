package songrepo

import "main/internal/models"

func (r *SongRepo) Create(group *models.Song) (uint, error) {
	err := r.db.Create(group).Error

	return group.ID, err
}
