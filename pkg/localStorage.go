package pkg

import (
	"go_alert_bot/internal/service/events"
	"sync"

	"go_alert_bot/internal/entities"
)

func NewCounterMap() *CounterMap {
	return &CounterMap{
		m: make(map[events.Event]int),
	}
}

type CounterMap struct {
	mx sync.Mutex
	m  map[events.Event]int
}

func (c *CounterMap) Load(key events.Event) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *CounterMap) Store(key events.Event, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *CounterMap) DeleteKey(key events.Event) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.m, key)
}

func (c *CounterMap) GetMap() map[events.Event]int {
	return c.m
}

func NewStorageMap() *StorageMap {
	return &StorageMap{
		m: make(map[entities.ChannelLink]events.EventChan, 10),
	}
}

type StorageMap struct {
	mx sync.Mutex
	m  map[entities.ChannelLink]events.EventChan
}

func (ec *StorageMap) Load(key entities.ChannelLink) (events.EventChan, bool) {
	ec.mx.Lock()
	defer ec.mx.Unlock()
	val, ok := ec.m[key]
	return val, ok
}

func (ec *StorageMap) Store(key entities.ChannelLink, value events.EventChan) events.EventChan {
	ec.mx.Lock()
	defer ec.mx.Unlock()
	ec.m[key] = value
	return value
}

func (ec *StorageMap) GetMap() map[entities.ChannelLink]events.EventChan {
	return ec.m
}
