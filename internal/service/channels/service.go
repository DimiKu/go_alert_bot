package channels

import (
	"errors"
	"fmt"

	"go_alert_bot/internal"
	"go_alert_bot/internal/db_operations"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg/link_gen"
)

type ChannelRepo interface {
	CreateTelegramChannel(channel db_operations.ChannelDb) error
	IsExistChannel(channel db_operations.ChannelDb) bool
	CreateStdoutChannel(channel db_operations.ChannelDb) error
}

type ChannelService struct {
	storage ChannelRepo
}

func NewChannelService(storage ChannelRepo) *ChannelService {
	return &ChannelService{storage: storage}
}

func (chs *ChannelService) CreateChannel(channel internal.ChannelDto) (internal.ChannelLinkDto, error) {

	link := internal.ChannelLinkDto(link_gen.LinkGenerate())

	channelDb := db_operations.ChannelDb{
		UserId:       channel.UserId,
		TgChatId:     channel.TgChatId,
		ChannelLink:  db_operations.ChannelLink(link),
		ChannelType:  channel.ChatType,
		FormatString: channel.FormatString,
	}

	if !chs.storage.IsExistChannel(channelDb) {
		switch channelDb.ChannelType {
		case entities.TelegramChatType:
			err := chs.storage.CreateTelegramChannel(channelDb)
			if err != nil {
				return 0, fmt.Errorf("failed to create channel")
			}
		case entities.StdoutChatType:
			err := chs.storage.CreateStdoutChannel(channelDb)
			if err != nil {
				return 0, fmt.Errorf("failed to create channel")
			}

		}

		return link, nil

	} else {
		return 0, errors.New("channel already exist")
	}

	return 0, nil
}
