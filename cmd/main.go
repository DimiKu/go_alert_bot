package main

import (
	"fmt"
	"go_alert_bot/pkg/db_operations"
	h "go_alert_bot/pkg/handlers"
	"go_alert_bot/pkg/service/channels"
	"go_alert_bot/pkg/service/chats"
	"go_alert_bot/pkg/service/users"
	"net/http"
)

func CheckChannel(c chan string) {
	for {
		fmt.Printf("Event is %s", <-c)
	}
}

func main() {
	db := db_operations.NewDBAdminManage()

	storage := db_operations.NewStorage(db.DBCreate("alertsbot"))
	storage.CreateDatabase()
	eventChan := make(chan string)

	userService := users.NewUserService(storage)
	chatService := chats.NewChatService(storage)
	channelService := channels.NewChannelService(storage)

	http.HandleFunc("/event", h.CreateEventInChannelHandler(eventChan))
	http.HandleFunc("/create_user", h.NewUserHandleFunc(userService))
	http.HandleFunc("/create_chat", h.NewChatHandleFunc(chatService))
	http.HandleFunc("/create_channel", h.NewChannelHandleFunc(channelService))

	go CheckChannel(eventChan)

	http.ListenAndServe(":8081", nil)

}
