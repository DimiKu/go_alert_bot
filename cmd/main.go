package main

import (
	"go_alert_bot/pkg/db_operations"
	h "go_alert_bot/pkg/handlers"
	"net/http"
)

func main() {
	db := db_operations.NewDBAdminManage()
	Storage := db_operations.NewStorage(db.DBCreate("alertsbot"))
	Storage.CreateDatabase()

	http.HandleFunc("/event", h.CreateEventHandler)

	// TODO это нужно переделать
	// TODO описать тут интерфейс
	http.HandleFunc("/create_user", h.NewUserHandleFunc(Storage))

	http.ListenAndServe(":8081", nil)
}
