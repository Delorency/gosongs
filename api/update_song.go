package api

import (
	"encoding/json"
	d "main/database"
	schema "main/schema"
	v "main/validator"
	"net/http"
	"strconv"
)

// @Summary Update song
// @Tags songs
// @Accept json
// @Produce json
//
// @Param id path int true "Song id"
//
// @Param input body schema.UpdateSong true "Input data"
//
// @Success 200 {array} schema.Song
// @Failure 500 {string} string "Internal Server Error"
// @Router /songs/{id} [put]
func UpdateSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Введено не число", http.StatusBadRequest)
		return
	}
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

	err = s.UpdateSongDB(num, &song)

	if err != nil {
		http.Error(w, "Ошибка обновления данных", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}
