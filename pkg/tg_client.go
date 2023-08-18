package pkg

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramClient struct {
	bot *tgbotapi.BotAPI
}

func New(apiKey string) *TelegramClient {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		fmt.Errorf("failed to create client")
	}

	return &TelegramClient{
		bot: bot,
	}
}
func (c *TelegramClient) SendMessage(text string, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, err := c.bot.Send(msg)
	return err
}
