//go:generate mockgen -source service.go -destination service_mock.go -package events
package events

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"time"

	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/entities"
	"go_alert_bot/internal/service/dto"
)

var ErrChannelNotFound = errors.New("channel not exist")

type EventRepo interface {
	GetChannelFromChannelLink(link entities.ChannelLink) *db_actions.ChannelDb
	IsExistChannelByChannelLink(link db_actions.ChannelLink) bool
	GetTelegramChannelByChannelLink(channel *db_actions.ChannelDb) (*db_actions.ChannelDb, error)
	GetStdoutChannelByChannelLink(channel *db_actions.ChannelDb) (*db_actions.ChannelDb, error)
}

type SendEventRepo interface {
	Send(event Event, channel *db_actions.ChannelDb, counter int) error
}

type Event struct {
	Key string `json:"key"`
	// UserId int    `json:"user_id"`
	link entities.ChannelLink
}

type EventService struct {
	storage         EventRepo
	eventMap        *StorageMap
	eventCounterMap *CounterMap
	SendEventRepos  map[string]SendEventRepo
	l               *zap.Logger
}

type EventChan chan Event

func NewEventService(storage EventRepo,
	clientsList map[string]SendEventRepo,
	l *zap.Logger,
) *EventService {
	eventMap := NewStorageMap()
	eventCounter := NewCounterMap()
	return &EventService{
		storage:         storage,
		eventMap:        eventMap,
		eventCounterMap: eventCounter,
		SendEventRepos:  clientsList,
		l:               l,
	}

}

func (es *EventService) CreateNewChannel() EventChan {
	eventChannel := make(chan Event, 10)
	return eventChannel
}

func (es *EventService) AddEventInChannel(event dto.EventDto, channelLinkDto dto.ChannelLinkDto) (string, error) {
	var channelLinkToChannel entities.ChannelLink
	var eventToChannel Event

	if !es.storage.IsExistChannelByChannelLink(db_actions.ChannelLink(channelLinkDto)) {
		return "", ErrChannelNotFound
	}
	channelLinkToChannel = entities.ChannelLink(channelLinkDto)
	eventToChannel = Event{Key: event.Key, link: channelLinkToChannel}

	eventChan, ok := es.eventMap.Load(channelLinkToChannel)
	if !ok {
		es.l.Info("event was added", zap.String("event", event.Key))
		eventChan = es.eventMap.Store(channelLinkToChannel, es.CreateNewChannel())
	}
	eventChan <- eventToChannel

	return "Event added", nil
}

func (es *EventService) Send(event Event, channel *db_actions.ChannelDb, counter int) {
	client := es.SendEventRepos[channel.ChannelType]
	if err := client.Send(event, channel, counter); err != nil {
		es.l.Error("can't send message", zap.Error(err))
	}
}

func (es *EventService) CheckEventsInChan(ctx context.Context) error {
	for _, channel := range es.eventMap.GetMap() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Service is stopped")
				return nil
			case eventFromChannel := <-channel:
				eventCount, ok := es.eventCounterMap.Load(eventFromChannel)
				if !ok {
					es.eventCounterMap.Store(eventFromChannel, 0)
				}
				es.eventCounterMap.Store(eventFromChannel, eventCount+1)
			}
		}
	}
	return nil
}

func (es *EventService) SendMessagesFromMap(ctx context.Context) error {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			es.l.Info("ticker gone in SendMessagesFromMap")
			for event := range es.eventCounterMap.GetMap() {
				eventCount, _ := es.eventCounterMap.Load(event)

				channel := es.storage.GetChannelFromChannelLink(event.link)

				var err error

				if channel != nil {
					switch channel.ChannelType {
					case entities.TelegramChatType:
						channel, err = es.storage.GetTelegramChannelByChannelLink(channel)
						es.l.Info("ticker gone in GetTelegramChannelByChannelLink")
						if err != nil {
							return err
						}
					case entities.StdoutChatType:
						channel, err = es.storage.GetStdoutChannelByChannelLink(channel)
						if err != nil {
							return err
						}
					}
				}
				es.Send(event, channel, eventCount)
				es.eventCounterMap.DeleteKey(event)
				es.l.Info("delete key", zap.String("event", event.Key))
				ticker.Reset(10 * time.Second)
			}
		}
	}
}

func (es *EventService) RunCheckEventChannel(ctx context.Context) error {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			es.l.Info("ticker gone to CheckEventsInChan")
			if err := es.CheckEventsInChan(ctx); err != nil {
				return err
			}
			ticker.Reset(10 * time.Second)
			go func() {
				es.l.Info("ticker gone to SendMessagesFromMap")
				if err := es.SendMessagesFromMap(ctx); err != nil {
					es.l.Error("error, %w", zap.Error(err))
				}
			}()
		}
	}
}
