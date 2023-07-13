package users

import (
	"errors"
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/db_operations"
)

type UserRepo interface {
	CreateUser(user db_operations.UserDb) error // TODO посмотреть
	CheckIfExistUser(user db_operations.UserDb) bool
}

type UserService struct {
	storage UserRepo
}

func NewUserService(storage UserRepo) *UserService {
	return &UserService{storage: storage}
}

func (us *UserService) CreateUser(user pkg.UserDto) (error, int) {
	userDb := db_operations.UserDb{UserID: user.UserId, ChatId: user.ChatId}

	if us.storage.CheckIfExistUser(userDb) {
		fmt.Println(user.UserId)
		return errors.New("user already exist"), user.UserId // TODO сделать свою ошибку
	}
	err := us.storage.CreateUser(userDb)
	if err != nil {
		return fmt.Errorf("failed to create user %w", err), 0
	}
	return nil, user.UserId
}

//func (us *UserService) CheckUser(user pkg.UserDto) bool {
//	userDb := db_operations.UserDb{UserID: user.UserId, ChatId: user.ChatId}
//	result := us.storage.CheckIfExistUser(userDb)
//	return result
//}
