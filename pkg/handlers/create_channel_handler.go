package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/service/channels"
	"net/http"
)

func NewChannelHandleFunc(service *channels.ChannelService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var channel pkg.ChannelDto

		if r.Method == http.MethodPost {
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&channel)
			if err != nil {
				fmt.Fprintf(w, "err %s", err)
			}
			err, channelLink := service.CreateChannel(channel)
			if err != nil {
				fmt.Fprintf(w, "err %s", err)
			} else {
				fmt.Fprintf(w, "your chanellink is %d", channelLink)
			}
		}

	}
}
