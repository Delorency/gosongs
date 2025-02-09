package db

import (
	"log"
	schema "main/schema"
)

func (s *Storage) GetSongsDB() ([]schema.Song, error) {
	rows, err := s.db.Query(`SELECT id, "group", song, releaseDate, text, link FROM song`)
	if err != nil {
		log.Println("Ошибка получения данных ", err)
		return nil, nil
	}

	defer rows.Close()

	var songs []schema.Song

	for rows.Next() {
		var song schema.Song
		if err := rows.Scan(&song.Id, &song.Group, &song.Song, &song.Release_date, &song.Text, &song.Link); err != nil {
			log.Println("Ошибка при чтении строки ", err)
			continue
		}
		songs = append(songs, song)
	}

	return songs, nil
}
