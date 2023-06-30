package chats

import (
	"fmt"
	"go_alert_bot/pkg/db_operations"
	"go_alert_bot/pkg/handlers"
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

func (cs *ChatService) CreateChat(chat handlers.ChatDto) error {
	chatDb := db_operations.ChatDb{UserId: chat.UserId, ChatId: chat.ChatId}
	err := cs.storage.CreateChat(chatDb)
	if err != nil {
		fmt.Errorf("failed to create chat %w", err)
	}
	return nil
}
