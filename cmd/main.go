package main

import (
	"go_alert_bot/pkg/db_operations"
	h "go_alert_bot/pkg/handlers"
	"go_alert_bot/pkg/service/channels"
	"go_alert_bot/pkg/service/chats"
	"go_alert_bot/pkg/service/users"
	"net/http"
)

func main() {
	db := db_operations.NewDBAdminManage()

	storage := db_operations.NewStorage(db.DBCreate("alertsbot"))
	storage.CreateDatabase()

	userService := users.NewUserService(storage)
	chatService := chats.NewChatService(storage)
	channelService := channels.NewChannelService(storage)

	http.HandleFunc("/event", h.CreateEventHandler)
	http.HandleFunc("/create_user", h.NewUserHandleFunc(userService))
	http.HandleFunc("/create_chat", h.NewChatHandleFunc(chatService))
	http.HandleFunc("/create_channel", h.NewChannelHandleFunc(channelService))

	http.ListenAndServe(":8081", nil)
}
