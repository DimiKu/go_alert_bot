package db_operations

import (
	"fmt"

	"github.com/google/uuid"
)

type ChatUUID string

type TelegramChat struct {
	ChatUUID     string      `db:"chat_uuid"`
	UserId       int         `db:"user_id"`
	TgChatId     int64       `db:"telegram_chat_id"`
	FormatString string      `db:"format_string"`
	ChannelLink  ChannelLink `db:"channel_link"`
}

type StdoutChat struct {
	ChatUUID     uuid.UUID   `db:"chat_uuid"`
	UserId       int         `db:"user_id"`
	FormatString string      `db:"format_string"`
	ChannelLink  ChannelLink `db:"channel_link"`
}

func (s *Storage) CreateTelegramChatInDB(chat TelegramChat) (ChatUUID, error) {
	chat.ChatUUID = (uuid.New()).String()

	// TODO вот это стоит улучшить. Нужно иметь возможность указать много чатов для одного channel_link
	q := `INSERT INTO telegram_chats (user_id, chat_uuid, telegram_chat_id, format_string, channel_link) values ($1, $2, $3, $4, $5)`

	_, err := s.conn.Exec(q, chat.UserId, chat.ChatUUID, chat.TgChatId, chat.FormatString, chat.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create chat, %w", err)
	}

	chatUuid := ChatUUID(chat.ChatUUID)
	return chatUuid, err
}

func (s *Storage) CreateStdoutChatInDB(chat StdoutChat) (ChatUUID, error) {
	chat.ChatUUID = uuid.New()
	// TODO то же что и выше
	q := `INSERT INTO stdout_chats (user_id, chat_uuid, format_string, channel_link) values ($1, $2, $3, $4)`
	_, err := s.conn.Exec(q, chat.UserId, chat.ChatUUID, chat.FormatString, chat.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create chat, %w", err)
	}

	chatUuidQuery := `SELECT chat_uuid FROM stdout_chats WHERE channel_link=$1`

	row, err := s.conn.Query(chatUuidQuery, chat.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to get chatID after chat creating in DB, %w", err)
	}

	var ChatUuidFromDB ChatUUID

	for row.Next() {
		row.Scan(&ChatUuidFromDB)
	}

	return ChatUuidFromDB, err
}
