package events

import (
	"sync"

	"go_alert_bot/internal/entities"
)

func NewEventCounters() *EventCounters {
	return &EventCounters{
		m: make(map[Event]int),
	}
}

type EventCounters struct {
	mx sync.Mutex
	m  map[Event]int
}

func (c *EventCounters) Load(key Event) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *EventCounters) Store(key Event, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

// TODO можно ли возврашать ошибку?
func (c *EventCounters) DeleteKey(key Event) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.m, key)
}

func (c *EventCounters) GetMap() map[Event]int {
	return c.m
}

func NewEventChanStorage() *EventChanStorage {
	return &EventChanStorage{
		m: make(map[entities.ChannelLink]EventChan, 10),
	}
}

type EventChanStorage struct {
	mx sync.Mutex
	m  map[entities.ChannelLink]EventChan
}

func (ec *EventChanStorage) Load(key entities.ChannelLink) (EventChan, bool) {
	ec.mx.Lock()
	defer ec.mx.Unlock()
	val, ok := ec.m[key]
	return val, ok
}

func (ec *EventChanStorage) Store(key entities.ChannelLink, value EventChan) EventChan {
	ec.mx.Lock()
	defer ec.mx.Unlock()
	ec.m[key] = value
	return value
}

func (ec *EventChanStorage) GetMap() map[entities.ChannelLink]EventChan {
	return ec.m
}
