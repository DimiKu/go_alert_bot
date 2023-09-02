package events

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go_alert_bot/internal"
	"go_alert_bot/internal/db_operations"
	"go_alert_bot/internal/entities"
)

var ErrChannelNotFound = errors.New("channel not exist")

type EventRepo interface {
	GetChannelFromChannelLink(link entities.ChannelLink) db_operations.ChannelDb
	IsExistChannelByChannelLink(link internal.ChannelLinkDto) bool
}

type SendEventRepo interface {
	Send(event Event, channel db_operations.ChannelDb, counter int)
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

func (es *EventService) AddEventInChannel(event internal.EventDto, channelLinkDto internal.ChannelLinkDto) (string, error) {
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

// TODO подумать над тем чтобы не зависеть от слоя контроллеров
func (es *EventService) Send(event Event, channel db_operations.ChannelDb, counter int) {
	switch channel.ChannelType {
	case entities.TelegramChatType:
		client := es.SendEventRepos[entities.TelegramChatType]
		client.Send(event, channel, counter)
	case entities.StdoutChatType:
		client := es.SendEventRepos[entities.StdoutChatType]
		client.Send(event, channel, counter)
	}
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

				// TODO разбить логику функции
				channel := es.storage.GetChannelFromChannelLink(event.link)
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
			go es.SendMessagesFromMap(ctx)
		}

	}

}
