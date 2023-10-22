package db_actions

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/google/uuid"
)

type ChatUUID string

type TgChatIds pq.Int64Array

type TelegramChat struct {
	ChatUUID     string      `db:"chat_uuid"`
	UserId       int         `db:"user_id"`
	TgChatIds    TgChatIds   `db:"telegram_chat_id"`
	FormatString string      `db:"format_string"`
	ChannelLink  ChannelLink `db:"channel_link"`
}

type StdoutChat struct {
	ChatUUID     uuid.UUID   `db:"chat_uuid"`
	UserId       int         `db:"user_id"`
	FormatString string      `db:"format_string"`
	ChannelLink  ChannelLink `db:"channel_link"`
}

func (s *Storage) CreateTelegramChatInDB(chat TelegramChat) (*ChatUUID, error) {
	chat.ChatUUID = (uuid.New()).String()

	_, err := s.conn.Exec(insertTelegramChat,
		chat.UserId,
		chat.ChatUUID,
		pq.Array(chat.TgChatIds),
		chat.FormatString,
		chat.ChannelLink)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat, %w", err)
	}

	chatUuid := ChatUUID(chat.ChatUUID)
	return &chatUuid, err
}

func (s *Storage) CreateStdoutChatInDB(chat StdoutChat) (*ChatUUID, error) {
	chat.ChatUUID = uuid.New()

	_, err := s.conn.Exec(insertStdoutChat, chat.UserId, chat.ChatUUID, chat.FormatString, chat.ChannelLink)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat, %w", err)
	}

	row, err := s.conn.Query(selectChatUuid, chat.ChannelLink)
	if err != nil {
		s.l.Error("failed to get chatID after chat creating in DB", zap.Error(err))
	}

	var ChatUuidFromDB ChatUUID

	for row.Next() {
		if err := row.Scan(&ChatUuidFromDB); err != nil {
			return nil, err
		}
	}

	return &ChatUuidFromDB, err
}

func (s *Storage) AddNewChatToExistChannel(chat *TelegramChat, chatUUID uuid.UUID) error {
	if chat != nil {
		_, err := s.conn.Exec(updateTelegramChat, pq.Array(chat.TgChatIds), chatUUID)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("can't update chat. Chat is empty")
}
