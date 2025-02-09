package api

import (
	d "main/database"
	"net/http"
)

func DeleteSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	err := s.DeleteSongDB(id)

	if err != nil {
		http.Error(w, "Ошибка при удалении песни", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
