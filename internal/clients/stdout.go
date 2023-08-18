package clients

import (
	"fmt"

	"go_alert_bot/internal/service/events"
)

type StdoutClient struct {
	formatString string
}

func NewStdoutClient() *StdoutClient {
	formatString := "Event %s was %d times"
	return &StdoutClient{formatString: formatString}
}

func (s *StdoutClient) SendEvent(event events.Event, chatID int64, counter int) {
	fmt.Printf(s.formatString, event.Key, counter)
}
