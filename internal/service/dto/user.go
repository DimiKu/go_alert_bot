package dto

import "errors"

var EmptyUserError = errors.New("user can't be empty")

type UserDto struct {
	UserId int `json:"user_id"`
	ChatId int `json:"chat_id"`
}

func (u *UserDto) Validate() error {
	if u.UserId == 0 {
		return EmptyUserError
	}
	
	if u.ChatId == 0 {
		return EmptyChatID
	}

	return nil
}
