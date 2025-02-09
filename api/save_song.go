package api

import (
	"encoding/json"
	d "main/database"
	schema "main/schema"
	"net/http"

	v "main/validator"
)

func SaveSong(s *d.Storage, w http.ResponseWriter, r *http.Request) {
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
	songDetail.Release_date = "Test"
	songDetail.Text = "Test"
	songDetail.Link = "Test"

	_, err := s.SaveSongDB(songDetail.Group, songDetail.Song, songDetail.Release_date, songDetail.Text, songDetail.Link)

	if err != nil {
		http.Error(w, "Ошибка сохранения", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
