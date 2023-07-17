package channels

import (
	"errors"
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/db_operations"
	"go_alert_bot/pkg/utils"
)

type ChannelRepo interface {
	CreateChannel(channel db_operations.ChannelDb) (error, int64)
	CheckChannel(channel db_operations.ChannelDb) bool
}

type ChannelService struct {
	storage ChannelRepo
}

func NewChannelService(storage ChannelRepo) *ChannelService {
	return &ChannelService{storage: storage}
}

// TODO должны ли сервисы быть в pkg
func (chs *ChannelService) CreateChannel(channel pkg.ChannelDto) (error, int64) {

	if chs.CheckChannel(channel) {
		channelLink := utils.LinkGenerate()
		channelDb := db_operations.ChannelDb{UserId: channel.UserId, ChatId: channel.ChatId, ChannelLink: channelLink}
		fmt.Println(channelDb.ChatId, channelDb.ChannelLink)
		err, _ := chs.storage.CreateChannel(channelDb)
		if err != nil {
			return fmt.Errorf("failed to create channel"), 0
		}
		return nil, channelLink
	} else {
		return errors.New("channel already exist"), 0
	}
	return nil, 0
}

func (chs *ChannelService) CheckChannel(channel pkg.ChannelDto) bool {
	channelDb := db_operations.ChannelDb{UserId: channel.UserId, ChatId: channel.ChatId}
	result := chs.storage.CheckChannel(channelDb)
	return result
}
