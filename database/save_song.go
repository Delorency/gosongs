package db

import (
	"log"
)

func (s *Storage) SaveSongDB(group, song, releaseDate, text, link string) (int64, error) {
	stmt, err := s.db.Prepare(`INSERT INTO song("group", song, releaseDate, text, link) VALUES($1,$2,$3,$4,$5)`)
	if err != nil {
		log.Println("Ошибка подготовки запроса на создание записи ", err)
		return 0, err
	}

	_, err = stmt.Exec(group, song, releaseDate, text, link)
	if err != nil {
		log.Println("Ошибка записи в таблицу ", err)
		return 0, err
	}
	return 1, nil
}
