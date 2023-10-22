package db_actions

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

type ChannelLink int64

type ChannelDb struct {
	UserId       int           `db:"user_id"`
	ChatUUID     uuid.UUID     `db:"chat_uuid"`
	ChannelLink  ChannelLink   `db:"channel_link"`
	TgChatIds    pq.Int64Array `db:"telegram_chat_id"`
	ChannelType  string        `db:"channel_type"`
	FormatString string        `db:"format_string"`
}

func (s *Storage) CreateTelegramChannel(channel ChannelDb) error {
	tgChat := TelegramChat{UserId: channel.UserId, TgChatIds: TgChatIds(channel.TgChatIds), FormatString: channel.FormatString, ChannelLink: channel.ChannelLink}

	chatUUID, err := s.CreateTelegramChatInDB(tgChat)
	if err != nil {
		s.l.Error("failed to create telegram chat in db", zap.Error(err))
	}

	if err = s.createTelegramChannelInDB(channel, chatUUID); err != nil {
		s.l.Error("failed to create telegram channel in db", zap.Error(err))
	}

	return nil
}

func (s *Storage) CreateStdoutChannel(channel ChannelDb) error {
	stdChat := StdoutChat{UserId: channel.UserId, FormatString: channel.FormatString, ChannelLink: channel.ChannelLink}

	chatUuid, err := s.CreateStdoutChatInDB(stdChat)
	if err != nil {
		return fmt.Errorf("failed to create stdout chat in db, %w", err)
	}

	if err := s.createStdoutChannelInDB(channel, chatUuid); err != nil {
		return fmt.Errorf("failed to create stdout chat in db, %w", err)
	}

	return nil
}

func (s *Storage) IsExistChannel(channel ChannelDb) bool {
	var channelTest ChannelDb

	row, _ := s.conn.Query(isExistChannelByChannelLink, channel.ChannelLink)
	err := row.Scan(&channelTest)
	for row.Next() {
		if err != nil {
			s.l.Error("can't scan channelTest in check", zap.Error(err))
		}
	}

	if channelTest.ChannelLink == 0 {
		return false
	}

	for row.Next() {
		if err := row.Scan(&channelTest); err != nil {
			s.l.Error("can't scan channelTest in check sec", zap.Error(err))
		}
	}

	return true
}

func (s *Storage) IsExistChannelByChannelLink(link ChannelLink) bool {
	var channelTest ChannelDb

	row, err := s.conn.Query(selectChannelByChannelLink, link)
	if err != nil {
		s.l.Error("failed to select from channels, %w", zap.Error(err))
	}

	for row.Next() {
		if err := row.Scan(&channelTest.UserId, &channelTest.ChatUUID, &channelTest.ChannelType, &channelTest.ChannelLink); err != nil {
			s.l.Error("failed to scan, %w", zap.Error(err))
		}
	}

	if channelTest.ChannelLink == 0 {
		return false
	}

	for row.Next() {
		if err := row.Scan(&channelTest); err != nil {
			s.l.Error("can't scan channelTest in check by channel link", zap.Error(err))
		}
	}

	return true
}

func (s *Storage) createTelegramChannelInDB(channel ChannelDb, chatUuid *ChatUUID) error {
	if chatUuid != nil {
		_, err := s.conn.Exec(insertTelegramChannel, channel.UserId, chatUuid, channel.ChannelType, channel.ChannelLink)
		if err != nil {
			return fmt.Errorf("failed to create channel %w", err)
		}
	}
	return nil
}

func (s *Storage) createStdoutChannelInDB(channel ChannelDb, chatUuid *ChatUUID) error {
	if chatUuid != nil {
		_, err := s.conn.Exec(insertStdoutChannel, channel.UserId, &chatUuid, channel.ChannelType, channel.ChannelLink)
		if err != nil {
			return fmt.Errorf("failed to create channel %w", err)
		}
	}
	return nil
}

func (s *Storage) GetChannelByChannelLink(link *ChannelLink) (*ChannelDb, error) {
	if link != nil {
		row, err := s.conn.Query(selectChannelByChannelLink, link)
		if err != nil {
			s.l.Error("failed to select from channels", zap.Error(err))
		}

		var channel ChannelDb

		for row.Next() {
			if err := row.Scan(&channel.UserId, &channel.ChatUUID, &channel.ChannelType, &channel.ChannelLink); err != nil {
				return nil, fmt.Errorf("failed to scan, %w", err)
			}
		}

		if channel.ChannelLink == 0 {
			return nil, nil
		}

		for row.Next() {
			if err := row.Scan(&channel); err != nil {
				return nil, err
			}
		}

		return &channel, nil
	}

	return nil, fmt.Errorf("link is epmty")
}

func (s *Storage) GetChatsByChatUUID(chatUUID *uuid.UUID) ([]int64, error) {
	if chatUUID != nil {
		var tgChats []int64
		row, err := s.conn.Query(selectChatsByChatUUID, chatUUID.String())
		if err != nil {
			return nil, errors.New("failed to select chats by chatUUID")
		}

		for row.Next() {
			if err := row.Scan(pq.Array(&tgChats)); err != nil {
				return nil, errors.New("failed to scan chats by chatUUID")
			}
		}

		return tgChats, nil
	}

	return nil, errors.New("can't find chat by chatUUID")
}
