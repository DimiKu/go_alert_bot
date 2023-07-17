package events

type EventRepo interface {
	AddEvent(event string) error
}

type Event struct {
	Key    string `json:"key"`
	UserId int    `json:"user_id"`
}

type DataChannel struct {
	Channel chan string
	// Receiver это channel link в будущем
	Receiver int64
}
type EventService struct {
	storage EventRepo
}
