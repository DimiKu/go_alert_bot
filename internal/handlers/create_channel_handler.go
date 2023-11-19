package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"go_alert_bot/internal/service/dto"
	"go_alert_bot/pkg/link_gen"
	"net/http"
	"strconv"

	"go_alert_bot/internal/service/channels"
)

// @Tags			channel
// @Router			/create_channel [post]
// @Description		Create channel
// @Param			RequestBody body dto.ChannelDto true "ChannelDto.go"
// @Produce			application/json
// @Success			200 {object} string "channel is "
func NewChannelHandleFunc(service *channels.ChannelService, l *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

				response := response{
					Status: true,
					Message: struct {
						Name  string `json:"name"`
						Value string `json:"value"`
					}{
						Name:  "channel",
						Value: strconv.FormatInt(int64(responseChannel.ChannelLink), 10),
					},
				}

				jsonRes, err := json.Marshal(response)
				if err != nil {
					l.Error("can't decode response", zap.Error(err))
				}

				if err := makeResponse(w, jsonRes); err != nil {
					l.Error("can't encode response", zap.Error(err))
				}
			}
		}
	}
}
