package api

import (
	d "main/database"
	"net/http"
	"strconv"
)

// @Summary Delete song
// @Tags songs
// @Accept json
// @Produce json
//
// @Param id path int true "Song id"
//
// @Success 204
// @Failure 500 {string} string "Internal Server Error"
// @Router /songs/{id} [delete]
func DeleteSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Введено не число", http.StatusBadRequest)
		return
	}
	err = s.DeleteSongDB(num)

	if err != nil {
		http.Error(w, "Ошибка при удалении песни", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
