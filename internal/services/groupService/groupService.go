package groupservice

import (
	groupdb "main/internal/database/groupDB"
	"main/internal/models"
)

type GroupServiceI interface {
	Create(*models.Group) error
	Update(id uint, data *models.Group) (*models.Group, error)
}

type groupService struct {
	repo groupdb.GroupDBI
}

func NewGroupService(repo groupdb.GroupDBI) GroupServiceI {
	return &groupService{repo}
}
