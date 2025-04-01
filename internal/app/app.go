package app

import (
	"fmt"
	"log"
	"main/internal/config"
	cont "main/internal/container"
	s "main/internal/transport/http"
	storage "main/storage"
	m "main/storage/migrations"
)

func Run() {
	cfg := config.MustLoad()

	db := storage.Psql(cfg.Db.Role, cfg.Db.Pass, cfg.Db.Name, cfg.Db.Host, cfg.Db.Port)

	container := cont.NewContainer(db)

	server := s.NewHTTPServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port, container)

	m.RunMigrations(db)

	log.Print(fmt.Sprintf("Starting server on %s", server.Addr))

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		panic("Must be implemented")
	}
}
