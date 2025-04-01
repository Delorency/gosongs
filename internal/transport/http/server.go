package server

import (
	"fmt"
	"net/http"

	"main/internal/container"
	r "main/internal/transport/http/router"
)

func NewHTTPServer(addr, port string, cont *container.Container) *http.Server {
	router := r.Router(cont)

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: router,
	}
}
