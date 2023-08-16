package entities

type Event struct {
	Key    string `json:"key"`
	UserId int    `json:"user_id"`
}

type EventQueue struct {
	Queue  []Event
	UserID int
}

type ChannelLink int64

func (ev EventQueue) NewEventQueue(UserId int) *EventQueue {
	return &EventQueue{UserID: UserId}
}

func AddNewEvent(queue EventQueue, event Event) EventQueue {
	queue.Queue = append(queue.Queue, event)
	return queue
}
