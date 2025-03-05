package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"main/internal/models"
)

func CreateSongTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240305_create_song_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Song{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(models.Song{}.TableName())
		},
	}
}
