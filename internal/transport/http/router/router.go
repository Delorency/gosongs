package router

import (
	"main/internal/container"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func AddMiddleware(router *chi.Mux) *chi.Mux {
	// router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	return router
}

func Router(cont *container.Container) http.Handler {
	router := AddMiddleware(chi.NewRouter())

	router.Mount("/groups", NewGroupRouter(cont))

	return router
}
