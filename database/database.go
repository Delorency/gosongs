package database

import (
	"database/sql"
	"log"
	"os"
)

type Storage struct {
	db *sql.DB
}

func initDB() (*Storage, error) {
	connStr := os.Getenv("connStr")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Connection error", err)
		return nil, err
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS song(
			id INTEGER PRIMARY KEY,
			group VARCHAR(255) NOT NULL,
			song TEXT NOT NULL,
			releaseDate VARCHAR(128) NOT NULL,
			text TEXT NOT NULL,
			link TEXT NOT NULL);
	`)
	if err != nil {
		log.Fatal("Preparing table initialization error", err)
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal("Table initialization error", err)
		return nil, err
	}

	return &Storage{db: db}, nil
}
func (s *Storage) SaveSong(group, song, releaseDate, text, link string) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO song(group, song, releaseDate, text, link) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Println("Preparing save song error", err)
		return 0, err
	}

	res, err := stmt.Exec(group, song, releaseDate, text, link)
	if err != nil {
		log.Println("Save song error", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("get last insert id error", err)
		return 0, err
	}

	return id, nil
}
func (s *Storage) GetSongs() (string, error)
