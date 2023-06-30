package entities

type Event struct {
	Key    string `json:"key"`
	UserId int    `json:"user_id"`
}

type EventQueue struct {
	Queue  []Event
	UserID int
}

func (ev EventQueue) NewEventQueue(UserId int) *EventQueue {
	// TODO создаем ли тут объект
	return &EventQueue{UserID: UserId}
}

func AddNewEvent(queue EventQueue, event Event) EventQueue {
	queue.Queue = append(queue.Queue, event)
	return queue
}

//func CollectMessages(event *Event, UserId int, queue *EventQueue) EventQueue{
//
//}
