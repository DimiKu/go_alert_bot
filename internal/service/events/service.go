package events

import (
	"context"
	"errors"
	"fmt"
	"go_alert_bot/internal/service/dto"
	"sync"
	"time"

	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/entities"
)

var ErrChannelNotFound = errors.New("channel not exist")

type EventRepo interface {
	GetChannelFromChannelLink(link entities.ChannelLink) *db_actions.ChannelDb
	IsExistChannelByChannelLink(link dto.ChannelLinkDto) bool
	GetTelegramChannelByChannelLink(channel *db_actions.ChannelDb) (*db_actions.ChannelDb, error)
	GetStdoutChannelByChannelLink(channel *db_actions.ChannelDb) (*db_actions.ChannelDb, error)
}

type SendEventRepo interface {
	Send(event Event, channel *db_actions.ChannelDb, counter int)
}

type Event struct {
	Key    string `json:"key"`
	UserId int    `json:"user_id"`
	link   entities.ChannelLink
}

type EventService struct {
	storage         EventRepo
	eventMap        *EventChanStorage
	eventCounterMap *EventCounters
	SendEventRepos  map[string]SendEventRepo
}

type EventChan chan Event

func NewEventService(storage EventRepo, clientsList map[string]SendEventRepo) *EventService {
	eventMap := NewEventChanStorage()
	eventCounter := NewEventCounters()
	return &EventService{storage: storage, eventMap: eventMap, eventCounterMap: eventCounter, SendEventRepos: clientsList}

}

func (es *EventService) CreateNewChannel() EventChan {
	eventChannel := make(chan Event, 10)
	return eventChannel
}

func (es *EventService) AddEventInChannel(event dto.EventDto, channelLinkDto dto.ChannelLinkDto) (string, error) {
	var channelLinkToChannel entities.ChannelLink
	var eventToChannel Event
	if !es.storage.IsExistChannelByChannelLink(channelLinkDto) {
		return "", ErrChannelNotFound
	}
	channelLinkToChannel = entities.ChannelLink(channelLinkDto)
	eventToChannel = Event{Key: event.Key, UserId: event.UserId, link: channelLinkToChannel}

	eventChan, ok := es.eventMap.Load(channelLinkToChannel)
	if !ok {
		eventChan = es.eventMap.Store(channelLinkToChannel, es.CreateNewChannel())
	}
	eventChan <- eventToChannel

	return "Event added", nil
}

func (es *EventService) Send(event Event, channel *db_actions.ChannelDb, counter int) {
	client := es.SendEventRepos[channel.ChannelType]
	client.Send(event, channel, counter)
}

func (es *EventService) CheckEventsInChan(ctx context.Context) error {

	for link, channel := range es.eventMap.GetMap() {
		fmt.Println(link, channel)
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
	timer := time.NewTimer(10 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-timer.C:
			for event := range es.eventCounterMap.GetMap() {
				eventCount, _ := es.eventCounterMap.Load(event)

				channel := es.storage.GetChannelFromChannelLink(event.link)

				var err error

				if channel != nil {
					switch channel.ChannelType {
					case entities.TelegramChatType:
						channel, err = es.storage.GetTelegramChannelByChannelLink(channel)
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
			}

			return nil
		}
	}
}

func (es *EventService) RunCheckEventChannel(ctx context.Context, wg *sync.WaitGroup) error {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := es.CheckEventsInChan(ctx); err != nil {
				return err
			}
			// TODO подумать чтобы отправлялось после ctx.Done
			ticker.Reset(5 * time.Second)
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				if err := es.SendMessagesFromMap(ctx); err != nil {
					fmt.Errorf("error, %w", err)
				}
			}(wg)
		}
	}
}
