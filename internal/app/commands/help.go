package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CreateButton() [][]tgbotapi.KeyboardButton {

	newButton := [][]tgbotapi.KeyboardButton{
		{{Text: "/help"}},
		{{Text: "/list"}},
	}
	return newButton
}
func (c *Commander) HelpCommand(inputMessage *tgbotapi.Message) {
	keyBrd := tgbotapi.ReplyKeyboardMarkup{Keyboard: CreateButton()}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - Get some help\n/list - Get list of all products")
	msg.ReplyMarkup = &keyBrd
	v, ok := msg.ReplyMarkup.(*tgbotapi.ReplyKeyboardMarkup)
	v.Selective = true
	fmt.Println("**************** ", v, ok)

	c.bot.Send(msg)

}

func init() {
	registeredCommands["help"] = (*Commander).HelpCommand
}
