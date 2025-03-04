package router

import (
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

func Router() http.Handler {
	router := chi.NewRouter()

	router = AddMiddleware(router)

	// Endpoints
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	return router
}
