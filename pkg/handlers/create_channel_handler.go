package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/service/channels"
	"go_alert_bot/pkg/utils"
	"net/http"
)

func NewChannelHandleFunc(service *channels.ChannelService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var channel pkg.ChannelDto

			err := json.NewDecoder(r.Body).Decode(&channel)
			if err != nil {
				fmt.Errorf("Failed to decode")
			}
			if service.CheckChannel(channel) {
				channel.ChannelLink = utils.LinkGenerate()
				service.CreateChannel(channel)
				fmt.Fprintf(w, "Your channel link is %d", channel.ChannelLink)
			} else {
				fmt.Fprintf(w, "Channel already exist")
			}
		}
	}
}
