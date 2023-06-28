package entities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id     int `db:"id"`
	ChatId int `db:"chat"`
}

var UserCounter int

func (u *User) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello, world"))
	var user User
	if r.Method == http.MethodPost {
		UserCounter += 1
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Errorf("Failed to create user")
		}
		user.Id = UserCounter

	}
}
