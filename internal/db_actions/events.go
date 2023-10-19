package db_actions

import (
	"fmt"
	"go_alert_bot/internal/entities"
)

func (s *Storage) GetChannelFromChannelLink(link entities.ChannelLink) *ChannelDb {
	var existChannel ChannelDb

	if err := s.conn.QueryRow(selectChannelByChannelLink, link).Scan(
		&existChannel.UserId,
		&existChannel.ChatUUID,
		&existChannel.ChannelType,
		&existChannel.ChannelLink,
	); err != nil {
		fmt.Errorf("failed to scan channel, %w", err)
	}

	return &existChannel
}

// TODO если несколько чатов. Наверно нужно использовать слайс
func (s *Storage) GetTelegramChannelByChannelLink(channel *ChannelDb) (*ChannelDb, error) {
	if err := s.conn.QueryRow(selectTelegramChat,
		channel.ChatUUID,
	).Scan(
		&channel.TgChatIds,
		&channel.FormatString,
	); err != nil {
		return nil, fmt.Errorf("failed to scan channel, %w", err)
	}

	return channel, nil
}

func (s *Storage) GetStdoutChannelByChannelLink(channel *ChannelDb) (*ChannelDb, error) {
	if err := s.conn.QueryRow(selectFormatStringByStdoutChat,
		channel.ChatUUID,
	).Scan(
		&channel.FormatString,
	); err != nil {
		return nil, fmt.Errorf("failed to scan channel, %w", err)
	}

	return channel, nil
}
