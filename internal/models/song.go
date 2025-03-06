package models

import (
	"time"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	ReleaseDate time.Time `gorm:"not null"`
	Text        string    `gorm:"not null"`
	Link        string    `gorm:"not null"`

	GroupID uint  `gorm:"not null"`
	Group   Group `gorm:"foreignKey:GroupID"`
}

func (Song) TableName() string {
	return "songs"
}
