package main

import (
	"log"
	"time"

	"github.com/Murkadymov/vsbot/internal/app/commands"

	"github.com/Murkadymov/vsbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(commands.GetToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	var productService *product.Service = product.NewService()
	var commander *commands.Commander = commands.NewCommander(bot, productService)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		commander.GetUpdate(update)

	}
}
