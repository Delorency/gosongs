package groupdb

import (
	"main/internal/models"
	"main/internal/schemes"

	"gorm.io/gorm"
)

type GroupDBI interface {
	List(group *models.Group, p *schemes.Pagination) ([]models.Group, error)
	Create(group *models.Group) error
	Update(id uint, data *models.Group) (*models.Group, error)
	Retrieve(group *models.Group) (models.Group, error)
}

type groupDB struct {
	db *gorm.DB
}

func NewGroupDB(db *gorm.DB) GroupDBI {
	return &groupDB{db}
}
