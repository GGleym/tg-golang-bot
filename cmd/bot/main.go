package main

import (
	"github/GGleym/telegram-todo-app-golang/internal/bot"
	"github/GGleym/telegram-todo-app-golang/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := config.InitConfig()
	initiatedBot, err := bot.InitBot(config.Token)

	if err != nil {
		log.Printf("Could not initiate the bot: %v", err)
	}

	initiatedBot.API.Debug = true

	updateConfig := bot.UpdateBot(30)
	updates := initiatedBot.API.GetUpdatesChan(updateConfig)
	
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.ReplyToMessageID = update.Message.MessageID

		
		if _, err := initiatedBot.API.Send(msg); err != nil {
            panic(err)
        }
	}
}
