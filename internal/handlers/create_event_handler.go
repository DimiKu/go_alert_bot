package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go_alert_bot/internal/service/events"
)

var UserCounter int

func CreateEventInChannelHandler(service *events.EventService, l *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("Event accepted!"))
		vars := mux.Vars(r)
		ChannelLink := vars["channelLink"]
		cnannelLinkInInt, err := strconv.ParseInt(ChannelLink, 10, 64)
		if err != nil {
			l.Error("Failed to parse channel link", zap.Error(err))
		}
		ChannelLinkDto := dto.ChannelLinkDto(cnannelLinkInInt)
		var event dto.EventDto
		if r.Method == http.MethodPost {
			err := json.NewDecoder(r.Body).Decode(&event)
			if err != nil {
				l.Error("Failed to decode event", zap.Error(err))
			}

			res, err := service.AddEventInChannel(event, ChannelLinkDto)
			if errors.Is(err, events.ErrChannelNotFound) {
				fmt.Fprintf(w, "Channel not exist")
			}
			if err != nil {
				l.Error("Failed to decode", zap.Error(err))
				return
			}
			if res != "" {
				fmt.Fprintf(w, "Event is %s", event.Key)
			}

		}
	}
}
