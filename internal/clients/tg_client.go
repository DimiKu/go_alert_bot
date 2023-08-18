package clients

import (
	"strconv"
	"strings"

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

func (es *TelegramClient) SendEvent(
	event events.Event,
	chatId int64,
	counter int,
) {
	// fmt.Printf("\nEvent  %s was %d times sended to link %s", event.Key, counter, link)
	counterStr := strconv.Itoa(counter)
	msg := strings.Join([]string{"Event", event.Key, " was ", counterStr, " times"}, " ")
	es.tgClient.SendMessage(msg, chatId)
}
