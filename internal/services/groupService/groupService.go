package groupservice

import (
	groupdb "main/internal/database/groupDB"
	"main/internal/models"
	"main/internal/schemes"
)

type GroupServiceI interface {
	List(*schemes.Pagination) (*[]models.Group, error)
	Create(*models.Group) error
	Update(uint, *models.Group) (*models.Group, error)
	Retrieve(uint) (*models.Group, error)
}

type groupService struct {
	repo groupdb.GroupDBI
}

func NewGroupService(repo groupdb.GroupDBI) GroupServiceI {
	return &groupService{repo}
}
