package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func InitDB() (*Storage, error) {
	godotenv.Load()

	conn := os.Getenv("conn")

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Ошибка соединения ", err)
		return nil, err
	}

	log.Println("Успешное соединение с БД")

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS song(
			id SERIAL PRIMARY KEY,
			"group" VARCHAR(255) NOT NULL,
			song TEXT NOT NULL,
			releaseDate VARCHAR(128) NOT NULL,
			text TEXT NOT NULL,
			link TEXT NOT NULL);
	`)
	if err != nil {
		log.Fatal("Ошибка подготовки инициализации таблиц ", err)
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("Ошибка инициализации таблиц ", err)
		return nil, err
	}
	log.Println("Успешная инициализация таблиц")

	return &Storage{db: db}, nil
}
func (s *Storage) SaveSong(group, song, releaseDate, text, link string) (int64, error) {
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
func (s *Storage) GetSongs() (string, error) {
	return "", nil
}
