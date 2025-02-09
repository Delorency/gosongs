package api

import (
	"encoding/json"
	d "main/database"
	"main/schema"
	"net/http"
)

func GetSongs(song *schema.FilterPag, pg *schema.TextPagination, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	songs, err := s.GetSongsDB(song, pg)

	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}
