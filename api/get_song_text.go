package api

import (
	"database/sql"
	"encoding/json"
	d "main/database"
	schema "main/schema"
	"net/http"
	"strings"
)

func GetSongText(pg schema.Pagination, id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	if pg.Skip < 0 || pg.Limit < 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{})
		return
	}

	text, err := s.GetSongTextDB(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, "Ошибка при извлечении песни", http.StatusInternalServerError)
		}
		return
	}

	// Разделение по \n\n
	arr := strings.Split(text, "\n\n")
	start := pg.Skip * pg.Limit
	end := start + pg.Limit

	if start >= len(arr) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{})
		return
	}

	if end > len(arr) {
		end = len(arr)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arr[start:end])
}
