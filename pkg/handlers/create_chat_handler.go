package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/service/chats"
	"net/http"
)

func NewChatHandleFunc(service *chats.ChatService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var chat pkg.ChatDto

			err := json.NewDecoder(r.Body).Decode(&chat)
			if err != nil {
				fmt.Errorf("Failed to decode")
			}

			service.CreateChat(chat)
		}
	}
}
