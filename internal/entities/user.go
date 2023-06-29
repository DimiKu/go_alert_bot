package entities

type User struct {
	Id     int `json:"id"`
	ChatId int `json:"chat_id"`
}

func (u *User) NewUser(id, ChatId int) User {
	return User{id, ChatId}

}
