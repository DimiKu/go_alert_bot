package events

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"go_alert_bot/internal"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg"
)

var ErrChannelNotFound = errors.New("channel not exist")

type EventRepo interface {
	GetChatsFromChannelLink(link entities.ChannelLink) int64
	IsExistChannelByChannelLink(link internal.ChannelLinkDto) bool
}

type SendEventRepo interface {
	SendEvent(event Event, chatID int64, counter int)
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
	SendEventRepo   SendEventRepo
}

type EventChan chan Event

func NewEventService(storage EventRepo, client SendEventRepo) *EventService {
	eventMap := NewEventChanStorage()
	eventCounter := NewEventCounters()
	return &EventService{storage: storage, eventMap: eventMap, eventCounterMap: eventCounter, SendEventRepo: client}

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

func (es *EventService) SendEvent(event Event, counter int, link entities.ChannelLink, client *pkg.TelegramClient) {
	// fmt.Printf("\nEvent  %s was %d times sended to link %s", event.Key, counter, link)
	chatId := es.storage.GetChatsFromChannelLink(link)
	counterStr := strconv.Itoa(counter)
	msg := strings.Join([]string{"Event", event.Key, " was ", counterStr, " times"}, " ")
	client.SendMessage(msg, chatId)
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
	//ticker := time.NewTicker(10 * time.Second)
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-timer.C:
			for event := range es.eventCounterMap.GetMap() {
				eventCount, _ := es.eventCounterMap.Load(event)
				chatID := es.storage.GetChatsFromChannelLink(event.link)
				es.SendEventRepo.SendEvent(event, chatID, eventCount)
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
