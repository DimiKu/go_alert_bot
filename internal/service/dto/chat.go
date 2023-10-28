package dto

import (
	"errors"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg"
	"strconv"
)

var EmptyChatID = errors.New("chat_id can't be empty")

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

	if c.UserId == 0 {
		return EmptyUserError
	}

	if c.ChatType != entities.StdoutChatType && c.ChatType != entities.TelegramChatType {
		return InvalidChatType
	}

	if c.ChatType == entities.TelegramChatType || c.ChatType == "" {
		_, err := pkg.ConvertStrToInt64Slice(c.TgChatId)
		if err != nil {
			return InvalidChatId
		}
	}

	return nil
}
