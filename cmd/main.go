package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go_alert_bot/internal/db_operations"
	"go_alert_bot/internal/handlers"
	"go_alert_bot/internal/service/channels"
	"go_alert_bot/internal/service/chats"
	"go_alert_bot/internal/service/events"
	"go_alert_bot/internal/service/users"
	"net/http"
	"sync"
)

func main() {
	db := db_operations.NewDBAdminManage()

	storage := db_operations.NewStorage(db.DBCreate("alertsbot"))
	storage.CreateDatabase()
	// ctx, cancel := context.WithCancel(context.Background())
	// ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	userService := users.NewUserService(storage)
	chatService := chats.NewChatService(storage)
	channelService := channels.NewChannelService(storage)
	eventService := events.NewEventService(storage)
	eventChannel := eventService.CreateNewChannel()

	router := mux.NewRouter()
	router.HandleFunc("/event/{channelLink}", handlers.CreateEventInChannelHandler(eventService, eventChannel))
	router.HandleFunc("/create_user", handlers.NewUserHandleFunc(userService))
	router.HandleFunc("/create_chat", handlers.NewChatHandleFunc(chatService))
	router.HandleFunc("/create_channel", handlers.NewChannelHandleFunc(channelService))

	go func() {
		err := eventService.RunCheckEventChannel(ctx, wg)
		cancel()
		if err != nil {
			fmt.Printf("error %w", err)
		}
	}()

	http.ListenAndServe(":8081", router)

}
