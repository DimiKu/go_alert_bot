package clients

import (
	"fmt"

	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/service/events"
)

type StdoutClient struct {
	formatString string
}

func NewStdoutClient() *StdoutClient {
	formatString := "Event %s was %d times, "
	return &StdoutClient{formatString: formatString}
}

func (s *StdoutClient) Send(event events.Event, channel *db_actions.ChannelDb, counter int) error {
	returnedString := s.formatString + channel.FormatString
	fmt.Printf(returnedString, event.Key, counter)

	return nil
}
