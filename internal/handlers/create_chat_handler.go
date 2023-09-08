package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go_alert_bot/internal"
	"go_alert_bot/internal/service/chats"
)

// TODO пока чаты хочу убрать. Они не нужны, если мы регистрируем их при регистрации channel. В будущем планирую
// TODO использовать их для расширения отправки по channel_link

func NewChatHandleFunc(service *chats.ChatService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var chat internal.ChatDto

			err := json.NewDecoder(r.Body).Decode(&chat)
			if err != nil {
				fmt.Errorf("Failed to decode")
			}

			if err := service.CreateChat(chat); err != nil {
				fmt.Fprintf(w, "failed to create chat %s", err)
			}
		}
	}
}
