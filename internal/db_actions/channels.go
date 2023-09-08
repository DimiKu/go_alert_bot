package db_actions

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"go_alert_bot/internal"
)

type ChannelLink int64

type ChannelDb struct {
	UserId       int           `db:"user_id"`
	ChatUUID     uuid.UUID     `db:"chat_uuid"`
	ChannelLink  ChannelLink   `db:"channel_link"`
	TgChatIds    pq.Int64Array `db:"telegram_chat_id"`
	ChannelType  string        `db:"channel_type"`
	FormatString string        `db:"format_string"`
}

func (s *Storage) CreateTelegramChannel(channel ChannelDb) error {
	tgChat := TelegramChat{UserId: channel.UserId, TgChatIds: TgChatIds(channel.TgChatIds), FormatString: channel.FormatString, ChannelLink: channel.ChannelLink}

	chatUUID, err := s.CreateTelegramChatInDB(tgChat)
	if err != nil {
		fmt.Errorf("failed to create telegram chat in db, %w", err)
	}

	if err = s.createTelegramChannelInDB(channel, chatUUID); err != nil {
		fmt.Errorf("failed to create telegram channel in db, %w", err)
	}

	return nil
}

func (s *Storage) CreateStdoutChannel(channel ChannelDb) error {
	stdChat := StdoutChat{UserId: channel.UserId, FormatString: channel.FormatString, ChannelLink: channel.ChannelLink}

	chatUuid, err := s.CreateStdoutChatInDB(stdChat)
	if err != nil {
		return fmt.Errorf("failed to create stdout chat in db, %w", err)
	}

	if err := s.createStdoutChannelInDB(channel, chatUuid); err != nil {
		return fmt.Errorf("failed to create stdout chat in db, %w", err)
	}

	return nil
}

func (s *Storage) IsExistChannel(channel ChannelDb) bool {
	var channelTest ChannelDb

	row, _ := s.conn.Query(isExistChannelByChannelLink, channel.ChannelLink)
	row.Scan(&channelTest)

	if channelTest.ChannelLink == 0 {
		return false
	}

	for row.Next() {
		row.Scan(&channelTest)
	}

	return true
}

func (s *Storage) IsExistChannelByChannelLink(link internal.ChannelLinkDto) bool {
	var channelTest ChannelDb

	row, err := s.conn.Query(selectChannelByChannelLink, link)
	if err != nil {
		fmt.Errorf("failed to select from channels, %w", err)
	}

	for row.Next() {
		if err := row.Scan(&channelTest.UserId, &channelTest.ChatUUID, &channelTest.ChannelType, &channelTest.ChannelLink); err != nil {
			fmt.Errorf("failed to scan, %w", err)
		}
	}

	if channelTest.ChannelLink == 0 {
		return false
	}

	for row.Next() {
		row.Scan(&channelTest)
	}

	return true
}

func (s *Storage) createTelegramChannelInDB(channel ChannelDb, chatUuid *ChatUUID) error {
	if chatUuid != nil {
		_, err := s.conn.Exec(insertTelegramChannel, channel.UserId, chatUuid, channel.ChannelType, channel.ChannelLink)
		if err != nil {
			return fmt.Errorf("failed to create channel %w", err)
		}
	}
	return nil
}

func (s *Storage) createStdoutChannelInDB(channel ChannelDb, chatUuid *ChatUUID) error {
	if chatUuid != nil {
		_, err := s.conn.Exec(insertStdoutChannel, channel.UserId, &chatUuid, channel.ChannelType, channel.ChannelLink)
		if err != nil {
			return fmt.Errorf("failed to create channel %w", err)
		}
	}
	return nil
}
