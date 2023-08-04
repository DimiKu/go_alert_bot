package channels

import (
	"errors"
	"fmt"
	"go_alert_bot/internal"
	"go_alert_bot/internal/db_operations"
	"go_alert_bot/internal/utils"
)

type ChannelRepo interface {
	CreateChannel(channel db_operations.ChannelDb) (utils.ChannelLink, error)
	IsExistChannel(channel db_operations.ChannelDb) bool
}

type ChannelService struct {
	storage ChannelRepo
}

func NewChannelService(storage ChannelRepo) *ChannelService {
	return &ChannelService{storage: storage}
}

func (chs *ChannelService) CreateChannel(channel internal.ChannelDto) (utils.ChannelLink, error) {
	channelLink := utils.LinkGenerate() // TODO лучше возвращать инт64 и перекладывать или сразу возвращать нужный тип? Кажется тут смешиваются слои

	link := internal.ChannelLinkDto(channelLink)

	channelDb := db_operations.ChannelDb{UserId: channel.UserId, ChatId: channel.ChatId, ChannelLink: link}
	if !chs.storage.IsExistChannel(channelDb) {

		_, err := chs.storage.CreateChannel(channelDb)
		if err != nil {
			return 0, fmt.Errorf("failed to create channel")
		}
		return channelLink, nil
	} else {
		return 0, errors.New("channel already exist")
	}
	return 0, nil
}
