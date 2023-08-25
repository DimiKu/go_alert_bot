package internal

type ChannelDto struct {
	UserId       int    `json:"user_id"`
	TgChatId     int64  `json:"telegram_chat_id"`
	FormatString string `json:"format_string"`
	ChannelLink  ChannelLinkDto
	ChatType     string `json:"channel_type"`
}

type ChatDto struct {
	UserId       int    `json:"user_id"`
	TgChatId     int64  `json:"telegram_chat_id"`
	ChatType     string `json:"chat_type"`
	FormatString string `json:"format_string"`
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
