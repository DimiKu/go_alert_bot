package custom_errors

import (
	"errors"
)

var (
	UserAlreadyExist   = errors.New("user already exist")
	FailedToCreateUser = errors.New("failed add new user")
	UserNotExist       = errors.New("user not exist")
)
