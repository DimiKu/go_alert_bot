package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"go_alert_bot/pkg/link_gen"
	"io"
	"net/http"

	"go_alert_bot/internal/service/channels"
)

func NewChannelHandleFunc(service *channels.ChannelService, l *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var channel dto.ChannelDto

		if r.Method == http.MethodPost {
			if err := decode(r.Body, &channel); err != nil {
				l.Error("can't decode body from req", zap.Error(err))
			}

			if err := channel.Validate(); err != nil {
				l.Error("can't validate channel", zap.Error(err))
				if err := encode(w, channel.TgChatIds); err != nil {
					l.Error("can't encode channel.TgChatIds", zap.Error(err))
				}
			} else {

				link := dto.ChannelLinkDto(link_gen.LinkGenerate())
				channel.ChannelLink = link

				responseChannel, err := service.CreateChannel(channel)
				if err != nil {
					l.Error("can't create channel", zap.Error(err))
				}

				if err := encode(w, responseChannel); err != nil {
					l.Error("can't encode response", zap.Error(err))
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
