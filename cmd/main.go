package main

import (
	"go_alert_bot/pkg/db_operations"
	h "go_alert_bot/pkg/handlers"
	"net/http"
)

func main() {
	db := db_operations.NewDBAdminManage()
	storage := db_operations.NewStorage(db.DBCreate("alertsbot"))
	storage.CreateDatabase()

	http.HandleFunc("/", h.CreateEventHandler)

	http.ListenAndServe(":8081", nil)
}
