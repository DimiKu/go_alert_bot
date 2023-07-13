package db_operations

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserDb struct {
	UserID int `db:"user_id"`
	ChatId int `db:"chat_id"`
}

func (s *Storage) CreateUser(user UserDb) error {
	fmt.Println("Creating user")
	q := `INSERT INTO users (user_id, chat_id) values ($1, $2)`
	_, err := s.conn.Exec(q, user.UserID, user.ChatId)
	if err != nil {
		fmt.Errorf("Failed add new user")
	}
	return nil
}

func (s *Storage) CheckIfExistUser(user UserDb) bool {
	var checkUser UserDb
	userExist := true
	q := `SELECT user_id FROM users where user_id=$1`
	row := s.conn.QueryRow(q, user.UserID).Scan(&checkUser)
	fmt.Printf("user id is %d", user.UserID)
	// TODO посмотреть можно ли иначе
	if errors.Is(row, sql.ErrNoRows) {
		// TODO лучше ли так?
		userExist = false
	}
	return userExist
}
