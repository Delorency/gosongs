package api

import (
	"database/sql"
	"encoding/json"
	d "main/database"
	schema "main/schema"
	"net/http"
	"strings"
)

// @Summary Get song text
// @Tags songs
// @Accept json
// @Produce json
//
// @Param id path int true "Song id"
// @Param pagination query schema.Pagination false "Pagination"
//
// @Success 200 {array} []string
// @Failure 500 {string} string "Internal Server Error"
// @Router /songs/{id}/text [get]
func GetSongText(pg *schema.Pagination, id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	text, err := s.GetSongTextDB(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка при извлечении песни", http.StatusInternalServerError)
		}
		return
	}

	arr := strings.Split(text, "\n\n")

	start := pg.Offset * pg.Limit
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
