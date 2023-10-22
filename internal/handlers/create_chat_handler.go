package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"net/http"

	"go_alert_bot/internal/service/chats"
)

// TODO пока чаты хочу убрать. Они не нужны, если мы регистрируем их при регистрации channel. В будущем планирую
// TODO использовать их для расширения отправки по channel_link

func NewAddChatHandleFunc(service *chats.ChatService, l *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var chat dto.ChatDto

			err := json.NewDecoder(r.Body).Decode(&chat)
			if err != nil {
				l.Error("Failed to decode", zap.Error(err))
			}

			if err := service.AddChatToChannel(chat); err != nil {
				l.Error("failed to create chat", zap.Error(err))
			}
		}
	}
}
