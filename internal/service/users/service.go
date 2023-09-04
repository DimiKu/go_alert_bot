package users

import (
	"fmt"
	"go_alert_bot/internal"
	"go_alert_bot/internal/custom_errors"
	"go_alert_bot/internal/db_actions"
)

type UserRepo interface {
	CreateUser(user db_actions.UserDb) error
	CheckIfExistUser(user db_actions.UserDb) bool
}

type UserService struct {
	storage UserRepo
}

func NewUserService(storage UserRepo) *UserService {
	return &UserService{storage: storage}
}

func (us *UserService) CreateUser(user internal.UserDto) (int, error) {
	userDb := db_actions.UserDb{UserID: user.UserId, ChatId: user.ChatId}

	if us.storage.CheckIfExistUser(userDb) {
		return user.UserId, custom_errors.UserAlreadyExist
	}
	err := us.storage.CreateUser(userDb)
	if err != nil {
		return 0, fmt.Errorf("failed to create user %w", err)
	}
	return user.UserId, nil
}
