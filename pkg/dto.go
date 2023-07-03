package pkg

type ChannelDto struct {
	UserId      int `json:"user_id"`
	ChatId      int `json:"chat_id"`
	ChannelLink int64
}

type ChatDto struct {
	UserId int `json:"user_id"`
	ChatId int `json:"chat_id"`
}

type UserDto struct {
	Id     int `json:"user_id"`
	ChatId int `json:"chat_id"`
}
