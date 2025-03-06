package main

import (
	"fmt"
	"main/internal/config"
	s "main/internal/transport/http"
	storage "main/storage"
	m "main/storage/migrations"
)

func main() {
	cfg := config.MustLoad()

	server := s.NewHTTPServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	db := storage.Psql(cfg.Db.Role, cfg.Db.Pass, cfg.Db.Name, cfg.Db.Host, cfg.Db.Port)

	m.RunMigrations(db)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		panic("Must be implemented")
	}
}
