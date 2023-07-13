package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/service/users"
	"net/http"
)

func NewUserHandleFunc(service *users.UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var user pkg.UserDto

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Errorf("Failed to decode")
			}

			err, userId := service.CreateUser(user)
			if err != nil {
				fmt.Fprintf(w, "error: %w", err)
			} else {
				fmt.Fprintf(w, "your user_id is : %d", userId)
			}
		}
	}
}
