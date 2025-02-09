package db

import (
	"log"
	schema "main/schema"
)

func (s *Storage) UpdateSongDB(id string, song *schema.Song) error {
	_, err := s.db.Exec(`UPDATE song SET "group"=$1, song=$2, releaseDate=$3, text=$4, link=$5 WHERE id=$6`,
		song.Group, song.Song, song.Releasedate, song.Text, song.Link, id)

	if err != nil {
		log.Println("Ошибка обновления записи ", err)
		return err
	}

	return nil
}
