package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/pkg"
	"net/http"
)

var UserCounter int

func CreateEventInChannelHandler(eventChannel chan string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Event accepted!"))

		var event pkg.EventDto
		if r.Method == http.MethodPost {
			err := json.NewDecoder(r.Body).Decode(&event)
			if err != nil {
				fmt.Errorf("%w", err)
			}
			eventChannel <- event.Key
			if err != nil {
				fmt.Errorf("Failed to decode")
			}
			fmt.Fprintf(w, " Event is %s", event.Key)
		}
	}
}
