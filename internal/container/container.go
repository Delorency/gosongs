package container

import (
	groupdb "main/internal/database/groupDB"
	groupservice "main/internal/services/groupService"

	"gorm.io/gorm"
)

type Container struct {
	// Репозитории
	GroupRepo groupdb.GroupDBI

	// Сервисы
	GroupService groupservice.GroupServiceI
}

func NewContainer(db *gorm.DB) *Container {
	grouprepo := groupdb.NewGroupDB(db)

	groupserv := groupservice.NewGroupService(grouprepo)

	return &Container{
		GroupRepo:    grouprepo,
		GroupService: groupserv,
	}
}
