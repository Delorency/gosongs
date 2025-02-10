package api

import (
	"encoding/json"
	d "main/database"
	"main/schema"
	_ "main/swagger"
	"net/http"
)

// GetSongs - пример обработчика для получения пользователей из базы данных
// @Summary Get all songs
// @Description Get all songs from darabase
// @Tags songs
// @Accept json
// @Produce json
//
// @Param offset query int false "Pagination offset (default 0)"
// @Param limit query int false "Number of users to return (default 10)"
//
// @Param group query string false "group"
// @Param song query string false "song"
// @Param releasedate query string false "release date"
// @Param text query string false "text"
// @Param link query string false "link"
//
// @Success 200 {array} []swagger.Output
// @Failure 500 {string} string "Internal Server Error"
// @Router /songs [get]
func GetSongs(song *schema.FilterPag, pg *schema.TextPagination, s *d.Storage, w http.ResponseWriter, r *http.Request) {
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
