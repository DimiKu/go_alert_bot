package db_operations

import (
	"fmt"

	"github.com/google/uuid"

	"go_alert_bot/internal"
)

type ChannelLink int64

type ChannelDb struct {
	UserId       int         `db:"user_id"`
	ChatUUID     uuid.UUID   `db:"chat_uuid"`
	ChannelLink  ChannelLink `db:"channel_link"`
	TgChatId     int64       `db:"telegram_chat_id"` // TODO потом можно поменять на массив
	ChannelType  string      `db:"channel_type"`
	FormatString string      `db:"format_string"`
}

func (s *Storage) CreateTelegramChannel(channel ChannelDb) error {
	tgChat := TelegramChat{UserId: channel.UserId, TgChatId: channel.TgChatId, FormatString: channel.FormatString, ChannelLink: channel.ChannelLink}

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
		fmt.Errorf("failed to create stdout chat in db, %w", err)
	}

	if err := s.createStdoutChannelInDB(channel, chatUuid); err != nil {
		fmt.Errorf("failed to create stdout chat in db, %w", err)
	}

	return nil
}

func (s *Storage) IsExistChannel(channel ChannelDb) bool {
	var channelTest ChannelDb
	q := `SELECT user_id, chat_uuid, channel_link FROM channels where channel_link=$1`
	row, _ := s.conn.Query(q, channel.ChannelLink)
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

	q := `SELECT user_id, chat_uuid, channel_type, channel_link FROM channels where channel_link=$1`
	row, err := s.conn.Query(q, link)
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

// TODO потом можно объединить
func (s *Storage) createTelegramChannelInDB(channel ChannelDb, chatUuid ChatUUID) error {
	q := `INSERT INTO channels (user_id, chat_uuid, channel_type, channel_link) values ($1, $2, $3, $4)`
	_, err := s.conn.Exec(q, channel.UserId, chatUuid, channel.ChannelType, channel.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create channel %w", err)
	}

	return nil
}

func (s *Storage) createStdoutChannelInDB(channel ChannelDb, chatUuid ChatUUID) error {
	q := `INSERT INTO channels (user_id, chat_uuid, channel_type, channel_link) values ($1, $2, $3, $4)`
	_, err := s.conn.Exec(q, channel.UserId, chatUuid, channel.ChannelType, channel.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create channel %w", err)
	}

	return nil
}
