package main

import (
	"go_alert_bot/pkg/db_operations"
	h "go_alert_bot/pkg/handlers"
	"net/http"
)

func main() {
	db := db_operations.NewDBManage("alertsbot")
	// TODO чтобы пересоздавать базу. Уточнить, почему не отрабатывает defer
	db.DropDatabase()

	db.CreateDatabase()
	db.CreateUserTable("users")

	http.HandleFunc("/", h.CreateEventHandler)

	http.ListenAndServe(":8081", nil)
}
