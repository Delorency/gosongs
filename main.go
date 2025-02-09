package main

import (
	"log"
	"net/http"
	"strconv"

	"main/api"
	db "main/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	schema "main/schema"
)

func main() {
	storage, err := db.InitDB()
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
		api.SaveSong(storage, w, r)
	})
	r.Get("/songs/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		api.RetrieveSong(id, storage, w, r)
	})
	r.Delete("/songs/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		api.DeleteSong(id, storage, w, r)
	})
	r.Put("/songs/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		api.UpdateSong(id, storage, w, r)
	})
	r.Get("/songs/{id}/text", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		skip, _ := strconv.Atoi(r.URL.Query().Get("skip"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		skip--

		var pg schema.Pagination
		pg.Skip = skip
		pg.Limit = limit

		api.GetSongText(pg, id, storage, w, r)

	})

	log.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", r)
}
