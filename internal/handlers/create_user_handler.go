package handlers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"go_alert_bot/internal/service/users"
	"net/http"
)

func NewUserHandleFunc(service *users.UserService, l *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var user dto.UserDto

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				l.Error("Failed to decode", zap.Error(err))
			}

			userId, err := service.CreateUser(user)
			if err != nil {
				l.Error("can't create user", zap.Error(err))
			} else {
				fmt.Fprintf(w, "your user_id is : %d", userId)
			}
		}
	}
}
