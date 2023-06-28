package entities

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/pkg/db_operations"
	"net/http"
)

type User struct {
	Id     int `db:"id"`
	ChatId int `db:"chat"`
}

var Conn *db_operations.Storage
var UserCounter int

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Creating user"))
	var user User
	if r.Method == http.MethodPost {
		UserCounter += 1
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Errorf("Failed to create user")
		}
		user.Id = UserCounter
		Conn.CreateNewUser(user.Id, user.ChatId)
	}
}
