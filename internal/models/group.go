package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name string `gorm:"size:256;uniqueIndex"`
}

func (Group) TableName() string {
	return "groups"
}
