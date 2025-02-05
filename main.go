package main

import (
	"log"
	"net/http"

	api "main/api"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/songs", api.GetSongs)
	r.Post("/songs", api.AddSong)

	log.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", r)
}
