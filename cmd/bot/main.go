package main

import (
	"github/GGleym/telegram-todo-app-golang/internal/bot"
	"log"
)

func main() {
	bot, err := bot.InitBot("")

	if err != nil {
		log.Printf("Could not initiate the bot: %v", err)
	}
}
