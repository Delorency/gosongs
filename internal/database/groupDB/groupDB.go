package groupdb

import (
	"main/internal/models"
	"main/internal/schemes"

	"gorm.io/gorm"
)

type GroupDBI interface {
	List(*schemes.Pagination) (*[]models.Group, error)
	Create(*models.Group) error
	Update(uint, *models.Group) (*models.Group, error)
	Retrieve(uint) (*models.Group, error)
}

type groupDB struct {
	db *gorm.DB
}

func NewGroupDB(db *gorm.DB) GroupDBI {
	return &groupDB{db}
}
