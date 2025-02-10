package api

import (
	"database/sql"
	"encoding/json"
	d "main/database"
	"net/http"
)

// @Summary Retrieve one song
// @Tags songs
// @Accept json
// @Produce json
//
// @Param id path int true "Song id"
//
// @Success 200 {array} schema.Song
// @Failure 500 {string} string "Internal Server Error"
// @Router /songs/{id} [get]
func RetrieveSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	song, err := s.RetrieveSongDB(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка при извлечении песни", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*song)
}
