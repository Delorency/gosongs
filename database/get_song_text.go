package db

import "log"

func (s *Storage) GetSongTextDB(id string) (string, error) {
	var text string
	err := s.db.QueryRow("SELECT text FROM song WHERE id = $1", id).Scan(&text)

	if err != nil {
		log.Println("Ошибка получения данных ", err)
		return "", err
	}

	return text, nil
}
