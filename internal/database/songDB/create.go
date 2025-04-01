package songdb

import "main/internal/models"

func (r *SongDB) Create(group *models.Song) (uint, error) {
	err := r.db.Create(group).Error

	return group.ID, err
}
