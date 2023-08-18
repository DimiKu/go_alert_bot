package db_operations

import (
	"fmt"

	"go_alert_bot/internal"
)

type ChannelDb struct {
	UserId      int                     `db:"user_id"`
	ChatId      int64                   `db:"chat_id"`
	ChannelLink internal.ChannelLinkDto `db:"channel_link"`
	ChannelType string                  `db:"channel_type"`
}

func (s *Storage) CreateTelegramChannel(channel ChannelDb) error {
	q := `INSERT INTO channels (user_id, chat_id, chat_type, channel_link) values ($1, $2, $3)`
	_, err := s.conn.Exec(q, channel.UserId, channel.ChatId, channel.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create channel %w", err)
	}

	return nil
}

// TODO можно объеденить в одну с той что выше. Но как принимать форматирование
func (s *Storage) CreateStdoutChannel(channel ChannelDb) error {
	q := `INSERT INTO channels (user_id, chat_id, channel_type, channel_link) values ($1, $2, $3, $4)`
	_, err := s.conn.Exec(q, channel.UserId, channel.ChatId, channel.ChannelType, channel.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create channel %w", err)
	}

	return nil
}

func (s *Storage) IsExistChannel(channel ChannelDb) bool {
	var channelTest ChannelDb
	q := `SELECT user_id, chat_id, channel_link FROM channels where channel_link=$1`
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

	q := `SELECT user_id, chat_id, channel_type, channel_link FROM channels where channel_link=$1`
	row, err := s.conn.Query(q, link)
	if err != nil {
		fmt.Errorf("failed to select from channels, %w", err)
	}
	for row.Next() {
		if err := row.Scan(&channelTest.UserId, &channelTest.ChatId, &channelTest.ChannelType, &channelTest.ChannelLink); err != nil {
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
