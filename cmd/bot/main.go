package main

import (
	"github/GGleym/telegram-todo-app-golang/internal/bot"
	"github/GGleym/telegram-todo-app-golang/internal/commands"
	"github/GGleym/telegram-todo-app-golang/internal/config"
	"github/GGleym/telegram-todo-app-golang/internal/db"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := config.InitConfig()
	initiatedBot, err := bot.InitBot(config.Token)

	if err != nil {
		log.Printf("Could not initiate the bot: %v", err)
	}

	dbInstance := db.InitDB()
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

		commands.HandleCommands(update, &msg, dbInstance)
		sendMessage(initiatedBot.API, msg)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
