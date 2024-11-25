package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	API *tgbotapi.BotAPI
}

func InitBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	return &Bot{API: api}, nil
}

func UpdateBot(timeout int) tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = timeout
	
	return updateConfig
}
