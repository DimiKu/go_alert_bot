package db_operations

import (
	"database/sql"
	"fmt"
)

type ChannelDb struct {
	UserId      int   `db:"user_id"`
	ChatId      int   `db:"chat_id"`
	ChannelLink int64 `db:"channel_id"`
}

func (s *Storage) CreateChannel(channel ChannelDb) error {
	fmt.Println("Create channel")
	fmt.Printf("Channel for create %s", channel)
	q := `INSERT INTO channels (user_id, chat_id, channel_link) values ($1, $2, $3)`
	_, err := s.conn.Exec(q, channel.UserId, channel.ChatId, channel.ChannelLink)
	if err != nil {
		fmt.Errorf("failed to create channel %w", err)
	}

	return nil
}

func (s *Storage) CheckChannel(channel ChannelDb) bool {
	q := `SELECT user_id, chat_id, channel_link FROM channels where chat_id=$1`
	row := s.conn.QueryRow(q, channel.ChatId).Scan()
	if row == sql.ErrNoRows {
		return true
	} else {
		return false
	}

}
