package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func InitBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	return &Bot{api}, nil
}
