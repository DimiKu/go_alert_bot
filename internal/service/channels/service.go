package channels

import (
	"errors"
	"fmt"
	"go_alert_bot/internal"
	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg/link_gen"
	"strconv"
	"strings"
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

func (chs *ChannelService) CreateChannel(channel internal.ChannelDto) (internal.ChannelLinkDto, error) {
	link := internal.ChannelLinkDto(link_gen.LinkGenerate())

	trimmed := strings.Trim(channel.TgChatIds, "[]")
	stringsSlice := strings.Split(trimmed, ", ")
	tgIds := make([]int64, len(stringsSlice))

	for i, s := range stringsSlice {
		tgIds[i], _ = strconv.ParseInt(s, 10, 64)
	}

	channelDb := db_actions.ChannelDb{
		UserId:       channel.UserId,
		TgChatIds:    tgIds,
		ChannelLink:  db_actions.ChannelLink(link),
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
}
