package db

import (
	"encoding/json"
	"log"
	schema "main/schema"
)

func (s *Storage) GetSongsDB(song *schema.Filter, pg *schema.Pagination) ([]schema.Song, error) {
	data, _ := json.Marshal(song)

	var songMap map[string]string
	json.Unmarshal(data, &songMap)

	query, args := buildSongsQuery(songMap, pg)
	rows, err := s.db.Query(query, args...)
	if err != nil {
		log.Println("Ошибка получения данных ", err)
		return nil, nil
	}

	defer rows.Close()

	var songs []schema.Song

	for rows.Next() {
		var song schema.Song
		if err := rows.Scan(&song.Id, &song.Group, &song.Song, &song.Releasedate, &song.Text, &song.Link); err != nil {
			log.Println("Ошибка при чтении строки ", err)
			continue
		}
		songs = append(songs, song)
	}

	return songs, nil
}
