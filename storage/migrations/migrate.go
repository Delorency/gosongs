package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		CreateGroupTable(),
		CreateSongTable(),
	})

	if err := m.Migrate(); err != nil {
		panic("Must be implemented")
	}
}
