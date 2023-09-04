package db_actions

import (
	"database/sql"
	"errors"
	"fmt"
	"go_alert_bot/internal/custom_errors"
)

type UserDb struct {
	UserID int `db:"user_id"`
	ChatId int `db:"chat_id"`
}

func (s *Storage) CreateUser(user UserDb) error {
	fmt.Println("Creating user")
	_, err := s.conn.Exec(insertUser, user.UserID, user.ChatId)
	if err != nil {
		return custom_errors.FailedToCreateUser
	}
	return nil
}

func (s *Storage) CheckIfExistUser(user UserDb) bool {
	var checkUser UserDb
	userExist := true
	row := s.conn.QueryRow(isExistUserByUserId, user.UserID).Scan(&checkUser)
	if errors.Is(row, sql.ErrNoRows) {
		userExist = false
	}
	return userExist
}
