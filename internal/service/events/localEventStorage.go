package events

import (
	"sync"

	"go_alert_bot/internal/entities"
)

func NewCounterMap() *CounterMap {
	return &CounterMap{
		m: make(map[Event]int),
	}
}

type CounterMap struct {
	mx sync.Mutex
	m  map[Event]int
}

func (c *CounterMap) Load(key Event) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *CounterMap) Store(key Event, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *CounterMap) DeleteKey(key Event) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.m, key)
}

func (c *CounterMap) GetMap() map[Event]int {
	return c.m
}

func NewStorageMap() *StorageMap {
	return &StorageMap{
		m: make(map[entities.ChannelLink]EventChan, 10),
	}
}

type StorageMap struct {
	mx sync.Mutex
	m  map[entities.ChannelLink]EventChan
}

func (ec *StorageMap) Load(key entities.ChannelLink) (EventChan, bool) {
	ec.mx.Lock()
	defer ec.mx.Unlock()
	val, ok := ec.m[key]
	return val, ok
}

func (ec *StorageMap) Store(key entities.ChannelLink, value EventChan) EventChan {
	ec.mx.Lock()
	defer ec.mx.Unlock()
	ec.m[key] = value
	return value
}

func (ec *StorageMap) GetMap() map[entities.ChannelLink]EventChan {
	return ec.m
}
