package chats

import (
	"fmt"
	"go_alert_bot/internal"
	"go_alert_bot/internal/db_operations"
)

type ChatRepo interface {
	CreateChat(chat db_operations.ChatDb) error
}

type ChatService struct {
	storage ChatRepo
}

func NewChatService(storage ChatRepo) *ChatService {
	return &ChatService{storage: storage}
}

func (cs *ChatService) CreateChat(chat internal.ChatDto) error {
	chatDb := db_operations.ChatDb{UserId: chat.UserId, ChatId: chat.ChatId}
	err := cs.storage.CreateChat(chatDb)
	if err != nil {
		fmt.Errorf("failed to create chat %w", err)
	}
	return nil
}
