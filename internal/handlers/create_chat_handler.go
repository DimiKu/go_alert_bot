package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"net/http"

	"go_alert_bot/internal/service/chats"
)

// @Tags			chat
// @Router			/add_chat [post]
// @Description		Add chat to exits channel
// @Param			RequestBody body dto.ChatDto true "ChatDto.go"
// @Produce			application/json
// @Success			200 {object} string "chat was added "
func NewAddChatHandleFunc(service *chats.ChatService, l *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var chat dto.ChatDto
			response := response{}

			err := json.NewDecoder(r.Body).Decode(&chat)
			if err != nil {
				response.Message.Name = "error"
				response.Message.Value = "Failed to decode chat"
				l.Error("Failed to decode", zap.Error(err))
			}

			if err := service.AddChatToChannel(chat); err != nil {
				response.Message.Name = "error"
				response.Message.Value = "Failed to create chat"
				l.Error("failed to create chat", zap.Error(err))
			}

			jsonResp, err := json.Marshal(response)
			if err != nil {
				l.Error("failed to decode response", zap.Error(err))
			}

			if err := makeResponse(w, jsonResp); err != nil {
				l.Error("failed to send event", zap.Error(err))
			}
		}
	}
}
