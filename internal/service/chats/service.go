package chats

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"strconv"
	"strings"

	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/entities"
)

type ChatRepo interface {
	CreateTelegramChatInDB(chat db_actions.TelegramChat) (*db_actions.ChatUUID, error)
	CreateStdoutChatInDB(chat db_actions.StdoutChat) (*db_actions.ChatUUID, error)
	GetChannelByChannelLink(link *db_actions.ChannelLink) (*db_actions.ChannelDb, error)
	AddNewChatToExistChannel(chat *db_actions.TelegramChat, chatUUID uuid.UUID) error
	GetChatsByChatUUID(chatUUID *uuid.UUID) ([]int64, error)
}

type ChatService struct {
	storage ChatRepo
	l       *zap.Logger
}

func NewChatService(storage ChatRepo, log *zap.Logger) *ChatService {
	return &ChatService{storage: storage, l: log}
}

func (cs *ChatService) AddChatToChannel(chat dto.ChatDto) error {
	switch chat.ChatType {
	case entities.TelegramChatType:
		trimmed := strings.Trim(chat.TgChatId, "[]")
		stringsSlice := strings.Split(trimmed, ", ")
		tgIds := make([]int64, len(stringsSlice))

		for i, s := range stringsSlice {
			tgIds[i], _ = strconv.ParseInt(s, 10, 64)
		}

		chatDB := db_actions.TelegramChat{
			TgChatIds:    tgIds,
			UserId:       chat.UserId,
			FormatString: chat.FormatString,
			ChannelLink:  db_actions.ChannelLink(chat.ChannelLink),
		}
		// TODO изменить вот эту функцию
		existChannel, err := cs.storage.GetChannelByChannelLink(&chatDB.ChannelLink)
		if err != nil {
			return fmt.Errorf("failed to create telegram chat, %w", err)
		}

		if existChannel == nil {
			return fmt.Errorf("empty channel from db")
		}

		newChatToRegistr := []int64(chatDB.TgChatIds)
		chatsForUpdate, err := cs.storage.GetChatsByChatUUID(&existChannel.ChatUUID)
		if err != nil {
			cs.l.Error("can't select chats by chatUUID", zap.Error(err))
		}

		newChatArray := append(chatsForUpdate, newChatToRegistr...)

		chatDB.TgChatIds = newChatArray

		if err := cs.storage.AddNewChatToExistChannel(&chatDB, existChannel.ChatUUID); err != nil {
			return err
		}

	case entities.StdoutChatType:
		chatDB := db_actions.StdoutChat{
			UserId:       chat.UserId,
			FormatString: chat.FormatString,
			ChannelLink:  db_actions.ChannelLink(chat.ChannelLink),
		}

		_, err := cs.storage.CreateStdoutChatInDB(chatDB)
		if err != nil {
			return fmt.Errorf("failed to create stdout chat, %w", err)
		}
	}

	return nil
}
