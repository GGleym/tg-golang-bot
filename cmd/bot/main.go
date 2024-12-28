package main

import (
	"github/GGleym/telegram-todo-app-golang/internal/bot"
	"github/GGleym/telegram-todo-app-golang/internal/config"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config.LoadEnv()
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("TOKEN variable was not set")
	}

	telegramBot, err := bot.InitBot(token)
	if err != nil {
		log.Printf("Could not initiate the bot: %v", err)
	}
	telegramBot.API.Debug = true
	log.Printf("Authorized on account %v", telegramBot.API.Self.UserName)

	handleTelegramUpdates(telegramBot)
}

func handleTelegramUpdates(telegramBot *bot.Bot) {
	updateConfig := bot.UpdateBot(60)
	updates := telegramBot.API.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if !update.Message.IsCommand() {
			msg.Text = "Введите команду"
			sendMessage(telegramBot.API, msg)
			continue
		}

		bot.HandleCommands(update, &msg)
		sendMessage(telegramBot.API, msg)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
