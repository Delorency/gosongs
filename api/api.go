package api

import (
	"io"
	"net/http"
)

type Song struct {
	ID          int    `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func AddSong(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Ошибка при чтении тела запроса", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() // Закрытие r.Body после использования

	// Выводим прочитанное тело запроса
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Resource created"))
}

func GetSongs(w http.ResponseWriter, r *http.Request) {

}
