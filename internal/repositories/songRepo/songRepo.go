package songrepo

import "gorm.io/gorm"

type SongRepo struct {
	db *gorm.DB
}

func NewSongRepo(db *gorm.DB) *SongRepo {
	return &SongRepo{db}
}
