package groupservice

import (
	groupdb "main/internal/database/groupDB"
	"main/internal/models"
)

type GroupServiceI interface {
	Create(m *models.Group) error
}

type groupService struct {
	repo groupdb.GroupDBI
}

func NewGroupService(repo groupdb.GroupDBI) GroupServiceI {
	return &groupService{repo}
}
