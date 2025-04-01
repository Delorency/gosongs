package songdb

import "gorm.io/gorm"

type SongDB struct {
	db *gorm.DB
}

func NewSongRepo(db *gorm.DB) *SongDB {
	return &SongDB{db}
}
