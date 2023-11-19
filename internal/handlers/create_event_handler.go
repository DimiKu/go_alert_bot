package handlers

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go_alert_bot/internal/service/events"
)

var UserCounter int

// @Tags			event
// @Router			/event [post]
// @Description		Create event
// @Param			RequestBody body dto.EventDto true "event.go"
// @Produce			application/json
// @Success			200 {object} string "event is "
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
			response := response{}
			err := json.NewDecoder(r.Body).Decode(&event)
			if err != nil {
				response.Message.Name = "error"
				response.Message.Value = "Failed to decode event"
				l.Error("Failed to decode event", zap.Error(err))
			}

			res, err := service.AddEventInChannel(event, ChannelLinkDto)
			if errors.Is(err, events.ErrChannelNotFound) {
				response.Message.Name = "error"
				response.Message.Value = "Channel not exist"
				l.Error("Channel not exist", zap.Error(err))
			}
			if err != nil {
				l.Error("Failed to decode", zap.Error(err))
				return
			}

			if res != "" {
				response.Status = true
				response.Message.Name = "event"
				response.Message.Value = event.Key
			}

			jsonResp, err := json.Marshal(response)
			if err != nil {
				l.Error("failed to decode response", zap.Error(err))
			}

			if err := makeResponse(w, jsonResp); err != nil {
				l.Error("failed to send event", zap.Error(err))
			}
		}
	}
}
