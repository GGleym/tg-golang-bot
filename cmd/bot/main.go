package main

import (
	"github/GGleym/telegram-todo-app-golang/internal/bot"
	"github/GGleym/telegram-todo-app-golang/internal/commands"
	"github/GGleym/telegram-todo-app-golang/internal/config"
	"github/GGleym/telegram-todo-app-golang/internal/router"
	"log"
	"net/http"
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

	router := router.Router()

	go func() {
		log.Println("HTTP server listening on port 4000")
		if err := http.ListenAndServe(":4000", router); err != nil {
			log.Fatalf("Failed to start a server on port 4000: %v", err)
		}
	}()

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

		commands.HandleCommands(update, &msg)
		sendMessage(telegramBot.API, msg)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
