package connector

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Psql(dbrole, dbpass, dbname, dbhost, dbport string) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%s password=%s dbname=%s host=%v port=%s sslmode=disable TimeZone=Asia/Shanghai",
			dbrole, dbpass, dbname, dbhost, dbport),

		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Must be implemented")
	}

	return db
}
