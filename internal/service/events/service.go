package events

import (
	"context"
	"fmt"
	"go_alert_bot/internal"
	"go_alert_bot/internal/entities"
	"go_alert_bot/pkg"
	"strconv"
	"strings"
	"sync"
	"time"
)

type EventRepo interface {
	GetChatsFromChannelLink(link entities.ChannelLink) int64
}

var EventMap = make(map[entities.ChannelLink]EventChan, 10) // TODO мапы перенести в поля
var EventMapCounter = make(map[Event]int, 10)               // TODO посмотреть sync.map, либо посмотреть использование мапы в горутинах конкурент ацес го
var TgClient = pkg.New(entities.TgToken)                    // TODO инициализировать в мейне, добавить как поле структуры типа

type Event struct {
	Key    string `json:"key"`
	UserId int    `json:"user_id"`
	link   entities.ChannelLink
}

type EventService struct {
	storage EventRepo
}

type EventChan chan Event

func NewEventService(storage EventRepo) *EventService {
	return &EventService{storage: storage}

}

func (es *EventService) CreateNewChannel() EventChan {
	eventChannel := make(chan Event, 10)
	return eventChannel
}

func (es *EventService) AddEventInChannel(event internal.EventDto, channelLinkDto internal.ChannelLinkDto) string {
	var channelLinkToChannel entities.ChannelLink
	var eventToChannel Event

	channelLinkToChannel = entities.ChannelLink(channelLinkDto)
	eventToChannel = Event{Key: event.Key, UserId: event.UserId, link: channelLinkToChannel}

	_, ok := EventMap[channelLinkToChannel]
	if !ok {
		EventMap[channelLinkToChannel] = es.CreateNewChannel()
	}
	EventMap[channelLinkToChannel] <- eventToChannel

	return "Event added"
}

func (es *EventService) SendEvent(event Event, counter int, link entities.ChannelLink, client *pkg.Client) {
	// fmt.Printf("\nEvent  %s was %d times sended to link %s", event.Key, counter, link)
	chatId := es.storage.GetChatsFromChannelLink(link)
	counterStr := strconv.Itoa(counter)
	msg := strings.Join([]string{"Event", event.Key, " was ", counterStr, " times"}, " ")
	client.SendMessage(msg, chatId)
}

func (es *EventService) CheckEventsInChan(ctx context.Context) {

	for link, channel := range EventMap {
		fmt.Println(link, channel)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Service is stopped")
				break
			case eventFromChannel := <-channel:
				_, ok := EventMapCounter[eventFromChannel]
				if !ok {
					EventMapCounter[eventFromChannel] = 0
				}
				EventMapCounter[eventFromChannel] += 1

			}
		}

	}
}

func (es *EventService) SendMessagesFromMap(ctx context.Context, client *pkg.Client, wg *sync.WaitGroup) error {
	//ticker := time.NewTicker(10 * time.Second)
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-timer.C:
			for event := range EventMapCounter {
				es.SendEvent(event, EventMapCounter[event], event.link, client)
				delete(EventMapCounter, event)

			}
			//ticker.Reset(10 * time.Second)
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
		case <-ticker.C: // TODO вызывать функцию с логикой тикера
			es.CheckEventsInChan(ctx)

			ticker.Reset(5 * time.Second)
			go es.SendMessagesFromMap(ctx, TgClient, wg)
		}

	}

}
