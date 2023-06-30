package users

import (
	"fmt"
	"go_alert_bot/pkg/db_operations"
	"go_alert_bot/pkg/handlers"
)

type UserRepo interface {
	CreateUser(user db_operations.UserDb) error // TODO посмотреть
}

type UserService struct {
	storage UserRepo
}

func NewUserService(storage UserRepo) *UserService {
	return &UserService{storage: storage}
}

func (us *UserService) CreateUser(user handlers.UserDto) error {
	userDb := db_operations.UserDb{Id: user.Id, ChatId: user.ChatId}
	err := us.storage.CreateUser(userDb)
	if err != nil {
		return fmt.Errorf("failed to create user %w", err)
	}
	return nil
}
