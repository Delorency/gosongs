package main

import (
	"fmt"
	"main/internal/config"

	s "main/internal/transport/http"
)

func main() {
	cfg := config.MustLoad()

	server := s.NewHTTPServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		panic("Must be implemented")
	}
}
