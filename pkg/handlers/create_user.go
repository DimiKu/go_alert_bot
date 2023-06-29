package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg/db_operations"
	"net/http"
)

var Conn *db_operations.Storage
var UserCount int

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
	var user entities.User
	if r.Method == http.MethodPost {
		UserCount += 1
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Errorf("Failed to create user")
		}
		user.Id = UserCount
		Conn.CreateNewUser(user.Id, user.ChatId)
	}
}
