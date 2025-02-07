package main

import (
	"log"
	"net/http"

	"main/api"
	database "main/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	storage, err := database.InitDB()
	if err != nil {
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/songs", func(w http.ResponseWriter, r *http.Request) {
		api.GetSongs(storage, w, r)
	})
	r.Post("/songs", func(w http.ResponseWriter, r *http.Request) {
		api.AddSong(storage, w, r)
	})

	log.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", r)
}
