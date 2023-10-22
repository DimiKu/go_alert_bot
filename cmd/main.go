package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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
	logger, _ := zap.NewProduction()

	db := db_actions.NewDBAdminManage(logger)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	storage := db_actions.NewStorage(db.DBCreate("alertsbot"), logger)

	if err := storage.CreateBasicTables(); err != nil {
		logger.Error("failed to create db", zap.Error(err))
	}

	ctx, cancel := context.WithCancel(context.Background())

	clientsList := map[string]events.SendEventRepo{
		"telegram": clients.NewTelegramClient(TgToken),
		"stdout":   clients.NewStdoutClient(),
	}

	userService := users.NewUserService(storage)
	chatService := chats.NewChatService(storage, logger)
	channelService := channels.NewChannelService(storage)
	eventService := events.NewEventService(storage, clientsList, logger)

	router := mux.NewRouter()
	router.HandleFunc("/event/{channelLink}", handlers.CreateEventInChannelHandler(eventService, logger))
	router.HandleFunc("/create_user", handlers.NewUserHandleFunc(userService, logger))
	router.HandleFunc("/add_chat", handlers.NewAddChatHandleFunc(chatService, logger))
	router.HandleFunc("/create_channel", handlers.NewChannelHandleFunc(channelService, logger))

	var wg sync.WaitGroup
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := eventService.RunCheckEventChannel(ctx, wg)
		if err != nil {
			logger.Error("RunCheckEventChannel failed", zap.Error(err))
		}
	}(&wg)

	srv := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: router,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()

		shutSignal := <-sigChan
		fmt.Println(shutSignal)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("error in shutdown", zap.Error(err))
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		cancel()
		logger.Info("Server is stopped")
	}
}
