package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/internal/entities"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello, world"))
	var user entities.User
	if r.Method == http.MethodPost {
		UserCounter += 1
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Errorf("Failed to create user")
		}
		user.Id = UserCounter

	}

}
