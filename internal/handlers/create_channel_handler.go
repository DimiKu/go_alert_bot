package handlers

import (
	"encoding/json"
	"fmt"
	"go_alert_bot/internal/service/dto"
	"io"
	"net/http"

	"go_alert_bot/internal/service/channels"
)

func NewChannelHandleFunc(service *channels.ChannelService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var channel dto.ChannelDto

		if r.Method == http.MethodPost {
			if err := decode(r.Body, &channel); err != nil {
				fmt.Fprintf(w, "err, %w", err)
			}

			if err := channel.Validate(); err != nil {
				fmt.Fprintf(w, "err %s. Channel was\n", err)
				encode(w, channel.TgChatIds)
			} else {
				responseChannel, err := service.CreateChannel(channel)
				if err != nil {
					fmt.Fprintf(w, "err, %w", err)
				}

				if err := encode(w, responseChannel); err != nil {
					fmt.Fprintf(w, "err, %w", err)
				}
			}
		}
	}
}

func encode(w http.ResponseWriter, object any) error {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(object); err != nil {
		return err
	}

	return nil
}

func decode(r io.Reader, object any) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&object)
	if err != nil {
		return err
	}

	return nil
}
