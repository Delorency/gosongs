package db

import (
	"log"
	schema "main/schema"
)

func (s *Storage) RetrieveSongDB(id int) (*schema.Song, error) {
	var song schema.Song

	err := s.db.QueryRow(`SELECT id, "group", song, releaseDate, text, link FROM song where id=$1`, id).Scan(
		&song.Id, &song.Group, &song.Song, &song.Releasedate, &song.Text, &song.Link)

	if err != nil {
		log.Println("Ошибка получения данных ", err)
		return nil, err
	}

	return &song, nil
}
