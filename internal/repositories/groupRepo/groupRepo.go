package grouprepo

import "gorm.io/gorm"

type GroupRepo struct {
	db *gorm.DB
}

func NewSongRepo(db *gorm.DB) *GroupRepo {
	return &GroupRepo{db}
}
