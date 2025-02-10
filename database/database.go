package db

import (
	"database/sql"
	"fmt"
	"log"
	schema "main/schema"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func InitDB() (*Storage, error) {
	godotenv.Load()
	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

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

func buildSongsQuery(params map[string]string, pg *schema.Pagination) (string, []interface{}) {
	query := `SELECT id, "group", song, releasedate, text, link FROM song WHERE 1=1`
	var args []interface{}
	argID := 1

	filterFields := map[string]string{
		"group":       "ILIKE",
		"song":        "ILIKE",
		"releasedate": "=",
		"text":        "ILIKE",
		"link":        "ILIKE",
	}

	for key, operator := range filterFields {
		if value, exists := params[key]; exists && value != "" {
			query += fmt.Sprintf(` AND "%s" %s $%d`, key, operator, argID)
			if operator == "ILIKE" {
				value = "%" + value + "%"
			}
			args = append(args, value)
			argID++
		}
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, pg.Limit, pg.Offset)

	return query, args
}
