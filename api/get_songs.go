package api

import (
	"encoding/json"
	d "main/database"
	"net/http"
)

func GetSongs(s *d.Storage, w http.ResponseWriter, r *http.Request) {
	songs, err := s.GetSongsDB()

	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}
