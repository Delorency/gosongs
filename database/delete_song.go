package db

import "log"

func (s *Storage) DeleteSongDB(id string) error {
	_, err := s.db.Exec("DELETE FROM song where id=$1", id)

	if err != nil {
		log.Println("Ошибка при удалении ", err)
		return err
	}

	return nil
}
