package commands

import (
	"fmt"

	"github.com/Murkadymov/vsbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var registeredCommands = map[string]func(c *Commander, inputMessage *tgbotapi.Message){}

func (c *Commander) GetUpdate(update tgbotapi.Update) {
	registeredCommands["list"] = (*Commander).ListCommand
	val, ok := registeredCommands[update.Message.Command()]

	if ok {
		val(c, update.Message)
	} else {
		c.DefaultBehaviour(update.Message)
	}

	fmt.Println("ALL COMMANDS: ", registeredCommands, val, ok, update.Message.Command())
}

func NewMessageHere(chatID int64, text string, msgID int) tgbotapi.MessageConfig {
	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID: chatID,
		},
		Text: text,
	}
	msg.ReplyToMessageID = msgID

	return msg
}

type Commander struct { //General constructor to minimize paramaters in functions
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}
