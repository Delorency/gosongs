package api

import (
	"encoding/json"
	d "main/database"
	"main/schema"
	_ "main/swagger"
	"net/http"
)

// @Summary Get all songs
// @Tags songs
// @Accept json
// @Produce json
//
// @Param pagination query schema.Pagination false "Pagination"
// @Param filter query schema.Filter false "Filters"
//
// @Success 200 {array} []schema.Song
// @Failure 500 {string} string "Internal Server Error"
// @Router /songs [get]
func GetSongs(song *schema.Filter, pg *schema.Pagination, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	songs, err := s.GetSongsDB(song, pg)

	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
	}

	if songs == nil {
		songs = []schema.Song{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}
