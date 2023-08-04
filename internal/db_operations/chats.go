package db_operations

import "fmt"

type ChatDb struct {
	UserId int `db:"user_id"`
	ChatId int `db:"id"`
}

func (s *Storage) CreateChat(chat ChatDb) error {
	fmt.Println("Create chat")
	// TODO  вот это стоит улучшить. Нужно иметь возможность указать много чатов для одного channel_link
	q := `INSERT INTO chat (user_id, chat_id) values ($1, $2)`
	_, err := s.conn.Exec(q, chat.UserId, chat.ChatId)
	if err != nil {
		fmt.Errorf("Failed to create chat")
	}
	return nil
}
