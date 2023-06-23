package main

import (
	h "go_alert_bot/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", h.CreateEventHandler)
	http.ListenAndServe(":8081", nil)
}
