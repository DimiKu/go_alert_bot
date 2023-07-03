package channels

import (
	"fmt"
	"go_alert_bot/pkg"
	"go_alert_bot/pkg/db_operations"
)

type ChannelRepo interface {
	CreateChannel(channel db_operations.ChannelDb) error
	CheckChannel(channel db_operations.ChannelDb) bool
}

type ChannelService struct {
	storage ChannelRepo
}

func NewChannelService(storage ChannelRepo) *ChannelService {
	return &ChannelService{storage: storage}
}

func (chs *ChannelService) CreateChannel(channel pkg.ChannelDto) error {
	channelDb := db_operations.ChannelDb{UserId: channel.UserId, ChatId: channel.ChatId, ChannelLink: channel.ChannelLink}
	err := chs.storage.CreateChannel(channelDb)
	if err != nil {
		fmt.Errorf("Failed to create channel %w", err)
	}
	return nil
}

func (chs *ChannelService) CheckChannel(channel pkg.ChannelDto) bool {
	channelDb := db_operations.ChannelDb{UserId: channel.UserId, ChatId: channel.ChatId}
	result := chs.storage.CheckChannel(channelDb)
	return result
}
