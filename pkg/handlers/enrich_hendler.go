package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/pkg/service/users"
	"net/http"
)

type UserDto struct {
	Id     int
	ChatId int
}

func NewUserHandleFunc(service *users.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var user UserDto

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Errorf("Failed to decode")
			}

			service.CreateUser(user)
		}
	}
}
