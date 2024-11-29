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

	log.Printf("Authorized on account %v", initiatedBot.API.Self.UserName)

	updateConfig := bot.UpdateBot(60)
	updates := initiatedBot.API.GetUpdatesChan(updateConfig)
	
	for update := range updates {
		if update.Message == nil {
			continue
		}

	 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if !update.Message.IsCommand() {
			msg.Text = "Введите команду"

			initiatedBot.API.Send(msg)

			continue
		}

        switch update.Message.Command() {
        case "help":
            msg.Text = "У меня есть команды /sayhi и /status."
        case "sayhi":
            msg.Text = "Привет :)"
        case "status":
            msg.Text = "У меня все хорошо."
        default:
            msg.Text = "Такой команды у меня нет :("
        }

		if _, err := initiatedBot.API.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
