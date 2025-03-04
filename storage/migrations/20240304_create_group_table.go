package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"main/internal/models"
)

func CreateGroupTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240304_create_group_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Group{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
