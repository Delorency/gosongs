package api

import (
	"database/sql"
	"encoding/json"
	d "main/database"
	"net/http"
)

func RetrieveSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	song, err := s.RetrieveSongDB(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "Ошибка при извлечении песни", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*song)
}
