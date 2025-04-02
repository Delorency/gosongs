package router

import (
	"main/internal/container"

	gh "main/internal/transport/http/handlers/groupHandler"

	"github.com/go-chi/chi"
)

func NewGroupRouter(cont *container.Container) *chi.Mux {
	router := AddMiddleware(chi.NewRouter())
	handler := gh.NewGroupHandler(cont.GroupService)

	router.Get("/", handler.List)
	router.Get("/{id}", handler.Retireve)
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)

	return router
}
