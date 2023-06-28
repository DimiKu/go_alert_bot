package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg/db_operations"
	"net/http"
)

func NewHandleFunc(storage *db_operations.Storage) func(http.ResponseWriter, *http.Request) {
	fmt.Println("Im here")
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.User
		if r.Method == http.MethodPost {
			UserCounter += 1
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Errorf("Failed to create user")
			}
			user.Id = UserCounter
			fmt.Printf("user is %s %s", user.Id, user.ChatId)

			storage.CreateNewUser(user.Id, user.ChatId)
		}
	}

}
