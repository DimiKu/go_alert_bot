package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"go_alert_bot/internal/clients"
	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/handlers"
	"go_alert_bot/internal/service/channels"
	"go_alert_bot/internal/service/chats"
	"go_alert_bot/internal/service/events"
	"go_alert_bot/internal/service/users"
)

func main() {
	db := db_actions.NewDBAdminManage()

	storage := db_actions.NewStorage(db.DBCreate("alertsbot"))
	if err := storage.CreateDatabase(); err != nil {
		fmt.Errorf("failed to create db, %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	clientsList := map[string]events.SendEventRepo{
		"telegram": clients.NewTelegramClient(TgToken),
		"stdout":   clients.NewStdoutClient(),
	}

	userService := users.NewUserService(storage)
	chatService := chats.NewChatService(storage)
	channelService := channels.NewChannelService(storage)
	eventService := events.NewEventService(storage, clientsList)

	router := mux.NewRouter()
	router.HandleFunc("/event/{channelLink}", handlers.CreateEventInChannelHandler(eventService))
	router.HandleFunc("/create_user", handlers.NewUserHandleFunc(userService))
	router.HandleFunc("/create_chat", handlers.NewChatHandleFunc(chatService))
	router.HandleFunc("/create_channel", handlers.NewChannelHandleFunc(channelService))

	go func() {
		err := eventService.RunCheckEventChannel(ctx, wg)
		if err != nil {
			fmt.Printf("error %w", err)
		}
	}()
	// TODO обработка сигналов
	if err := http.ListenAndServe(":8081", router); err != nil {
		fmt.Errorf("error is, %w", err)
		cancel()
	}
}
