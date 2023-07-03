package users

import (
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/db_operations"
)

type UserRepo interface {
	CreateUser(user db_operations.UserDb) error // TODO посмотреть
	CheckUser(user db_operations.UserDb) bool
}

type UserService struct {
	storage UserRepo
}

func NewUserService(storage UserRepo) *UserService {
	return &UserService{storage: storage}
}

func (us *UserService) CreateUser(user pkg.UserDto) error {
	userDb := db_operations.UserDb{Id: user.Id, ChatId: user.ChatId}
	err := us.storage.CreateUser(userDb)
	if err != nil {
		return fmt.Errorf("failed to create user %w", err)
	}
	return nil
}

func (us *UserService) CheckUser(user pkg.UserDto) bool {
	userDb := db_operations.UserDb{Id: user.Id, ChatId: user.ChatId}
	result := us.storage.CheckUser(userDb)
	return result
}
