package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"go_alert_bot/internal/service/users"
	"net/http"
	"strconv"
)

// @Tags			user
// @Router			/create_user [post]
// @Summary			create_user
// @Description		create_user
// @Param			RequestBody body dto.UserDto true "UserDto.go"
// @Produce			application/json
// @Success			200 {object} string "your user_id is user_id"
func NewUserHandleFunc(service *users.UserService, l *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var user dto.UserDto
			response := response{}

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				response.Message.Name = "error"
				response.Message.Value = "Failed to decode"
				l.Error("Failed to decode", zap.Error(err))
			}

			userId, err := service.CreateUser(user)
			if err != nil {
				response.Message.Name = "error"
				response.Message.Value = "Failed to decode"
				l.Error("can't create user", zap.Error(err))
			}

			response.Status = true
			response.Message.Name = "user"
			response.Message.Value = strconv.Itoa(userId)

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
