package db_operations

import (
	"fmt"

	"go_alert_bot/internal"
	"go_alert_bot/pkg/link_gen"
)

type ChannelDb struct {
	UserId      int                     `db:"user_id"`
	ChatId      int64                   `db:"chat_id"`
	ChannelLink internal.ChannelLinkDto `db:"channel_link"`
}

func (s *Storage) CreateChannel(channel ChannelDb) (link_gen.ChannelLink, error) {
	q := `INSERT INTO channels (user_id, chat_id, channel_link) values ($1, $2, $3)`
	_, err := s.conn.Exec(q, channel.UserId, channel.ChatId, channel.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create channel %w", err)
	}

	return 0, nil
}

func (s *Storage) IsExistChannel(channel ChannelDb) bool {
	var channelTest ChannelDb
	q := `SELECT user_id, chat_id, channel_link FROM channels where channel_link=$1`
	row, _ := s.conn.Query(q, channel.ChannelLink)
	fmt.Println("row is", row)
	row.Scan(&channelTest)
	if channelTest.ChannelLink == 0 {
		return false
	}
	for row.Next() {
		row.Scan(&channelTest)
	}

	return true
}
