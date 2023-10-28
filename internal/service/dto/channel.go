package dto

import (
	"errors"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg"
)

var (
	InvalidChatType = errors.New("invalid chat type. Choice telegram or stdout type")
	InvalidChatId   = errors.New("invalid telegram chat id. Chat id must be able to convert to int64")
)

type ChannelLinkDto int64

type ChannelDto struct {
	UserId       int            `json:"user_id"`
	TgChatIds    string         `json:"telegram_chat_id"`
	FormatString string         `json:"format_string"`
	ChatType     string         `json:"channel_type"`
	ChannelLink  ChannelLinkDto `json:"channel_link"`
}

func (c *ChannelDto) Validate() error {
	if c.ChatType != entities.StdoutChatType && c.ChatType != entities.TelegramChatType {
		return InvalidChatType
	}

	if c.ChatType == entities.TelegramChatType || c.ChatType == "" {
		_, err := pkg.ConvertStrToInt64Slice(c.TgChatIds)
		if err != nil {
			return InvalidChatId
		}
	}

	if c.UserId == 0 {
		return EmptyUserError
	}

	return nil
}
