package groupservice

import (
	"main/internal/models"
	"main/internal/schemes"
	gs "main/internal/schemes/groupSchemes"
)

type GroupRepoI interface {
	Create(*models.Group) (uint, error)
	Get(*models.Group, *schemes.Pagination) ([]models.Group, error)
	Retrieve(*models.Group) (models.Group, error)
	Update(uint, gs.Update, *models.Group) error
}

type GroupService struct {
	R GroupRepoI
}
