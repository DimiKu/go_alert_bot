package main

import (
	"go_alert_bot/pkg/db_operations"
	h "go_alert_bot/pkg/handlers"
	"go_alert_bot/pkg/service/chats"
	"go_alert_bot/pkg/service/users"
	"net/http"
)

func main() {
	db := db_operations.NewDBAdminManage()

	storage := db_operations.NewStorage(db.DBCreate("alertsbot"))
	storage.CreateDatabase()

	userService := users.NewUserService(storage)
	http.HandleFunc("/event", h.CreateEventHandler)

	// TODO это нужно переделать
	// TODO описать тут интерфейс
	http.HandleFunc("/create_user", h.NewUserHandleFunc(userService))

	chatService := chats.NewChatService(storage)
	http.HandleFunc("/create_chet", h.NewChatHandleFunc(chatService))

	http.ListenAndServe(":8081", nil)
}
