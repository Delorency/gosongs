package database

import (
	"database/sql"
	"log"
	"os"

	schema "main/schema"

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
func (s *Storage) GetSongs() ([]schema.Song, error) {
	rows, err := s.db.Query(`SELECT id, "group", song, releaseDate, text, link FROM song`)
	if err != nil {
		log.Println("Ошибка получения данных ", err)
		return nil, nil
	}

	defer rows.Close()

	var songs []schema.Song

	for rows.Next() {
		var song schema.Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			log.Println("Ошибка при чтении строки ", err)
			continue
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (s *Storage) DeleteSong(id string) error {
	_, err := s.db.Exec("DELETE FROM song where id=$1", id)

	if err != nil {
		log.Println("Ошибка при удалении ", err)
		return err
	}

	return nil
}

func (s *Storage) UpdateSong(id string, song *schema.Song) error {
	_, err := s.db.Exec(`UPDATE song SET "group"=$1, song=$2, releaseDate=$3, text=$4, link=$5 WHERE id=$6`,
		song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, id)

	if err != nil {
		log.Println("Ошибка обновления записи ", err)
		return err
	}

	return nil
}

func (s *Storage) RetrieveSong(id string) (*schema.Song, error) {
	rows, err := s.db.Query(`SELECT id, "group", song, releaseDate, text, link FROM song where id=$1`, id)

	if err != nil {
		log.Println("Ошибка получения данных ", err)
		return nil, nil
	}

	defer rows.Close()

	var song schema.Song
	rows.Next()

	if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
		log.Println("Ошибка при чтении строки ", err)
	}

	return &song, nil
}
