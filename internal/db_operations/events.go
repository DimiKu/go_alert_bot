package db_operations

import (
	"fmt"

	"go_alert_bot/internal/entities"
)

// TODO получать отдельно линк, по нему находить чат в отдельной функции
// TODO switсh должен быть в сервисе
func (s *Storage) GetChannelFromChannelLink(link entities.ChannelLink) ChannelDb {
	var existChannel ChannelDb

	q := `SELECT chat_uuid, user_id, channel_type, channel_link FROM channels WHERE channel_link=$1`

	if err := s.conn.QueryRow(q, link).Scan(
		&existChannel.ChatUUID,
		&existChannel.UserId,
		&existChannel.ChannelType,
		&existChannel.ChannelLink,
	); err != nil {
		fmt.Errorf("failed to scan channel, %w", err)
	}

	// TODO унести в сервис
	switch existChannel.ChannelType {
	case entities.TelegramChatType:
		q2 := `SELECT telegram_chat_id, format_string FROM telegram_chats WHERE chat_uuid=$1`
		if err := s.conn.QueryRow(q2, existChannel.ChatUUID).Scan(&existChannel.TgChatId, &existChannel.FormatString); err != nil {
			fmt.Errorf("failed to scan channel, %w", err)
		}
	case entities.StdoutChatType:
		q2 := `SELECT format_string FROM stdout_chats WHERE chat_uuid=$1`
		if err := s.conn.QueryRow(q2, existChannel.ChatUUID).Scan(&existChannel.FormatString); err != nil {
			fmt.Errorf("failed to scan channel, %w", err)
		}
	}

	return existChannel
}
