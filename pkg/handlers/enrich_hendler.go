package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg/db_operations"
	"net/http"
)

var user entities.User

func NewUserHandleFunc(storage *db_operations.Storage) func(http.ResponseWriter, *http.Request) {
	fmt.Println("Im here")
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			UserCounter += 1

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Errorf("Failed to decode")
			}

			fmt.Println(storage)
			user.Id = UserCounter
			storage.CreateNewUser(user.Id, user.ChatId)
		}
	}

}
