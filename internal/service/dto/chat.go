package dto

import (
	"go_alert_bot/internal/entities"
	"strconv"
)

type ChatDto struct {
	UserId       int    `json:"user_id"`
	TgChatId     string `json:"telegram_chat_id"`
	ChatType     string `json:"chat_type"`
	FormatString string `json:"format_string"`
	ChannelLink  int64  `json:"channel_link"`
}

func (c *ChatDto) Validate() error {
	if c.ChatType == entities.TelegramChatType {
		_, err := strconv.ParseInt(c.TgChatId, 10, 64)
		if err != nil {
			return InvalidChatId
		}
	}

	return nil
}
