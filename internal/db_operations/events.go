package db_operations

import (
	"fmt"
	"go_alert_bot/internal/entities"
)

func (s *Storage) GetChatsFromChannelLink(link entities.ChannelLink) int64 {
	fmt.Println("getting chats")
	var existChannel ChannelDb

	//q := `INSERT INTO users (user_id, chat_id) values ($1, $2)`
	q := `SELECT * FROM channels WHERE channel_link=$1`

	channel := s.conn.QueryRow(q, link).Scan(&existChannel.UserId, &existChannel.ChatId, &existChannel.ChannelLink)
	fmt.Println(channel)
	return existChannel.ChatId // TODO пока один
}
