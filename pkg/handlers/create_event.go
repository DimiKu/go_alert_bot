package handlers

import (
	db "go_alert_bot/pkg/db_operations"
	"net/http"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world"))
	db.DatabaseQueryExec("Select 1")
}
