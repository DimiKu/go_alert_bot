package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/internal"
	"go_alert_bot/internal/service/channels"
	"net/http"
)

func NewChannelHandleFunc(service *channels.ChannelService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var channel internal.ChannelDto

		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&channel)
			if err != nil {
				fmt.Fprintf(w, "err %s", err)
			}
			channelLink, err := service.CreateChannel(channel)
			if err != nil {
				fmt.Fprintf(w, "err %s", err)
			} else {
				fmt.Fprintf(w, "your chanellink is %d", channelLink)
			}
		}

	}
}
