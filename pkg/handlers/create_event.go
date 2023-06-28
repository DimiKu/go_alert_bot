package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/internal/entities"
	"net/http"
)

var UserCounter int

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Event accepted!"))

	var e entities.Event
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			fmt.Errorf("Failed to decode")
		}
		fmt.Fprintf(w, " Event is %s", e.Key)
	}

}
