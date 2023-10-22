package clients

import (
	"strconv"
	"strings"

	"go_alert_bot/internal/db_actions"
	"go_alert_bot/internal/service/events"
	"go_alert_bot/pkg"
)

type TelegramClient struct {
	token    string
	tgClient *pkg.TelegramClient
}

func NewTelegramClient(token string) *TelegramClient {
	return &TelegramClient{
		token:    token,
		tgClient: pkg.New(token),
	}
}

func (es *TelegramClient) Send(
	event events.Event,
	channel *db_actions.ChannelDb,
	counter int,
) error {
	// fmt.Printf("\nEvent  %s was %d times sended to link %s", event.Key, counter, link)
	counterStr := strconv.Itoa(counter)
	msg := strings.Join([]string{"Event", event.Key, " was ", counterStr, " times, ", channel.FormatString}, " ")
	for _, chat := range channel.TgChatIds {
		if err := es.tgClient.SendMessage(msg, chat); err != nil {
			return err
		}
	}

	return nil
}
