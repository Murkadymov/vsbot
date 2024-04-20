package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) DefaultBehaviour(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Ты гей\n/help - Get some help\n/list - Get list of all products")

	msg.ReplyToMessageID = inputMessage.MessageID
	c.bot.Send(msg)
}
