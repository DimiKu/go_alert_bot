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
			if service.CheckUser(user) {
				service.CreateUser(user)
				fmt.Fprintf(w, "User_id is %d", user.Id)
			} else {
				fmt.Fprintf(w, "User with %d is already exists", user.Id)
			}
		}
	}
}
