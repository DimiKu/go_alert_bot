package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go_alert_bot/internal"
	"go_alert_bot/internal/service/events"
)

var UserCounter int

func CreateEventInChannelHandler(service *events.EventService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Event accepted!"))
		vars := mux.Vars(r)
		ChannelLink := vars["channelLink"] // TODO добавить проверку на существование линка в бд
		cnannelLinkInInt, err := strconv.ParseInt(ChannelLink, 10, 64)
		if err != nil {
			fmt.Errorf("Failed to parse channel link")
		}
		ChannelLinkDto := internal.ChannelLinkDto(cnannelLinkInInt)
		var event internal.EventDto
		if r.Method == http.MethodPost {
			err := json.NewDecoder(r.Body).Decode(&event)
			if err != nil {
				fmt.Errorf("%w", err)
			}

			service.AddEventInChannel(event, ChannelLinkDto)
			if err != nil {
				fmt.Errorf("Failed to decode")
			}
			fmt.Fprintf(w, " Event is %s", event.Key)
		}
	}
}
