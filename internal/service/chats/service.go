package chats

import (
	"fmt"
	"go_alert_bot/internal/service/dto"
	"strconv"
	"strings"

	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/entities"
)

type ChatRepo interface {
	CreateTelegramChatInDB(chat db_actions.TelegramChat) (*db_actions.ChatUUID, error)
	CreateStdoutChatInDB(chat db_actions.StdoutChat) (*db_actions.ChatUUID, error)
}

type ChatService struct {
	storage ChatRepo
}

func NewChatService(storage ChatRepo) *ChatService {
	return &ChatService{storage: storage}
}

func (cs *ChatService) CreateChat(chat dto.ChatDto) error {
	switch chat.ChatType {
	case entities.TelegramChatType:
		trimmed := strings.Trim(chat.TgChatId, "[]")
		stringsSlice := strings.Split(trimmed, ", ")
		tgIds := make([]int64, len(stringsSlice))

		for i, s := range stringsSlice {
			tgIds[i], _ = strconv.ParseInt(s, 10, 64)
		}

		chatDb := db_actions.TelegramChat{
			TgChatIds:    tgIds,
			UserId:       chat.UserId,
			FormatString: chat.FormatString,
		}

		_, err := cs.storage.CreateTelegramChatInDB(chatDb)
		if err != nil {
			return fmt.Errorf("failed to create telegram chat, %w", err)
		}
	case entities.StdoutChatType:
		chatDB := db_actions.StdoutChat{
			UserId:       chat.UserId,
			FormatString: chat.FormatString,
		}

		_, err := cs.storage.CreateStdoutChatInDB(chatDB)
		if err != nil {
			return fmt.Errorf("failed to create stdout chat, %w", err)
		}
	}

	return nil
}
