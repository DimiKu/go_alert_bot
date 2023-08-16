package internal

type ChannelDto struct {
	UserId      int   `json:"user_id"`
	ChatId      int64 `json:"chat_id"`
	ChannelLink ChannelLinkDto
}

type ChatDto struct {
	UserId int `json:"user_id"`
	ChatId int `json:"chat_id"`
}

type UserDto struct {
	UserId int `json:"user_id"`
	ChatId int `json:"chat_id"`
}

type EventDto struct {
	Key         string `json:"key"`
	UserId      int    `json:"user_id"`
	ChannelLink ChannelLinkDto
}

type ChannelLinkDto int64
