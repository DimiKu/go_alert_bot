package chats

import (
	"fmt"

	"go_alert_bot/internal"
	"go_alert_bot/internal/db_operations"
	"go_alert_bot/internal/entities"
)

type ChatRepo interface {
	CreateTelegramChatInDB(chat db_operations.TelegramChat) (db_operations.ChatUUID, error)
	CreateStdoutChatInDB(chat db_operations.StdoutChat) (db_operations.ChatUUID, error)
}

type ChatService struct {
	storage ChatRepo
}

func NewChatService(storage ChatRepo) *ChatService {
	return &ChatService{storage: storage}
}

func (cs *ChatService) CreateChat(chat internal.ChatDto) error {
	switch chat.ChatType {
	case entities.TelegramChatType:
		chatDb := db_operations.TelegramChat{TgChatId: chat.TgChatId, UserId: chat.UserId, FormatString: chat.FormatString}

		_, err := cs.storage.CreateTelegramChatInDB(chatDb)
		if err != nil {
			fmt.Errorf("failed to create telegram chat, %w", err)
		}
	case entities.StdoutChatType:
		chatDB := db_operations.StdoutChat{UserId: chat.UserId, FormatString: chat.FormatString}

		_, err := cs.storage.CreateStdoutChatInDB(chatDB)
		if err != nil {
			fmt.Errorf("failed to create stdout chat, %w", err)
		}

	}

	return nil
}
