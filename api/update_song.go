package api

import (
	"encoding/json"
	d "main/database"
	schema "main/schema"
	v "main/validator"
	"net/http"
)

func UpdateSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	var song schema.Song

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	song.Id = id

	if err := validate.Struct(song); err != nil {
		v.HandleValidationErrors(w, err)
		return
	}

	err := s.UpdateSongDB(id, &song)

	if err != nil {
		http.Error(w, "Ошибка обновления данных", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
