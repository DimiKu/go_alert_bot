package db_operations

import (
	"database/sql"
	"fmt"
)

type UserDb struct {
	Id     int `db:"id"`
	ChatId int `db:"chat_id"`
}

func (s *Storage) CreateUser(user UserDb) error {
	fmt.Println("Creating user")
	q := `INSERT INTO users (id, chat_id) values ($1, $2)`
	_, err := s.conn.Exec(q, user.Id, user.ChatId)
	if err != nil {
		fmt.Errorf("Failed add new user")
	}
	return nil
}

func (s *Storage) CheckUser(user UserDb) bool {
	q := `SELECT id FROM users where id=$1`
	row := s.conn.QueryRow(q, user.Id).Scan()
	fmt.Printf("row is %s", row)
	if row == sql.ErrNoRows {
		return true
	} else {
		return false
	}
}
