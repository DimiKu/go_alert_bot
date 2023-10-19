//go:generate mockgen -source service.go -destination service_mock.go -package channels
package channels

import (
	"errors"
	"fmt"
	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/entities"
	"go_alert_bot/internal/service/dto"
	"go_alert_bot/pkg"
)

type ChannelRepo interface {
	CreateTelegramChannel(channel db_actions.ChannelDb) error
	IsExistChannel(channel db_actions.ChannelDb) bool
	CreateStdoutChannel(channel db_actions.ChannelDb) error
}

type ChannelService struct {
	storage ChannelRepo
}

func NewChannelService(storage ChannelRepo) *ChannelService {
	return &ChannelService{storage: storage}
}

func (chs *ChannelService) CreateChannel(channel dto.ChannelDto) (*dto.ChannelDto, error) {
	tgIds, err := pkg.ConvertStrToInt64Slice(channel.TgChatIds)
	if err != nil {
		return nil, err
	}

	channelDb := db_actions.ChannelDb{
		UserId:       channel.UserId,
		TgChatIds:    tgIds,
		ChannelLink:  db_actions.ChannelLink(channel.ChannelLink),
		ChannelType:  channel.ChatType,
		FormatString: channel.FormatString,
	}

	if !chs.storage.IsExistChannel(channelDb) {
		switch channelDb.ChannelType {
		case entities.TelegramChatType:
			err := chs.storage.CreateTelegramChannel(channelDb)
			if err != nil {
				return nil, fmt.Errorf("failed to create channel")
			}
		case entities.StdoutChatType:
			err := chs.storage.CreateStdoutChannel(channelDb)
			if err != nil {
				return nil, fmt.Errorf("failed to create channel")
			}
		}

		return &channel, nil

	} else {
		return nil, errors.New("channel already exist")
	}
}
