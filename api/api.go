package api

import (
	"encoding/json"
	"net/http"

	d "main/database"

	"github.com/go-playground/validator"

	schema "main/schema"
	v "main/validator"
)

var validate = validator.New()

func AddSong(s *d.Storage, w http.ResponseWriter, r *http.Request) {
	var input struct {
		Group string `json:"group" validate:"required"`
		Song  string `json:"song" validate:"required"`
	}

	// Обратботка тела запроса
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(input); err != nil {
		v.HandleValidationErrors(w, err)
		return
	}

	// Получение данных из внешнего апи
	// apiURL := fmt.Sprintf("%s?group=%s&song=%s", os.Getenv("API_URL"), input.Group, input.Song)
	// resp, err := http.Get(apiURL)
	// if err != nil {
	// 	http.Error(w, "Ошибка запроса к внешнему API", http.StatusInternalServerError)
	// 	return
	// }
	// defer resp.Body.Close()

	var songDetail schema.Song
	// if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
	// 	http.Error(w, "Ошибка обработки ответа API", http.StatusInternalServerError)
	// 	return
	// }

	songDetail.Group = input.Group
	songDetail.Song = input.Song
	songDetail.ReleaseDate = "Test"
	songDetail.Text = "Test"
	songDetail.Link = "Test"

	_, err := s.SaveSong(songDetail.Group, songDetail.Song, songDetail.ReleaseDate, songDetail.Text, songDetail.Link)

	if err != nil {
		http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetSongs(s *d.Storage, w http.ResponseWriter, r *http.Request) {
	songs, err := s.GetSongs()

	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

func DeleteSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	err := s.DeleteSong(id)

	if err != nil {
		http.Error(w, "Ошибка при удалении песни", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	var song schema.Song

	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	song.ID = id

	err := s.UpdateSong(id, &song)

	if err != nil {
		http.Error(w, "Ошибка обновления данных", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func RetrieveSong(id string, s *d.Storage, w http.ResponseWriter, r *http.Request) {
	song, err := s.RetrieveSong(id)

	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*song)
}
