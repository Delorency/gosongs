package main

import (
	"log"
	"net/http"
	"strconv"

	"main/api"
	db "main/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "main/docs"
	schema "main/schema"
)

func main() {
	storage, err := db.InitDB()
	if err != nil {
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Get("/songs", func(w http.ResponseWriter, r *http.Request) {
		var filpg schema.Filter
		var pg schema.Pagination

		filpg.Group = r.URL.Query().Get("group")
		filpg.Song = r.URL.Query().Get("song")
		filpg.Releasedate = r.URL.Query().Get("release_date")
		filpg.Text = r.URL.Query().Get("text")
		filpg.Link = r.URL.Query().Get("link")

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))

		if err != nil || offset < 0 {
			offset = 0
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		pg.Offset = offset
		pg.Limit = limit

		api.GetSongs(&filpg, &pg, storage, w, r)
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
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))

		if err != nil || offset < 0 {
			offset = 0
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || limit <= 0 {
			limit = 10
		}

		var pg schema.Pagination
		pg.Offset = offset
		pg.Limit = limit

		api.GetSongText(&pg, id, storage, w, r)

	})

	log.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe(":8080", r)
}
