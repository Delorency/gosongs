package server

import (
	"fmt"
	"net/http"

	r "main/internal/transport/http/router"
)

func NewHTTPServer(addr, port string) *http.Server {
	router := r.Router()

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", addr, port),
		Handler: router,
	}
}
