package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) ListCommand(inputMessage *tgbotapi.Message) {
	outputMsg := ""
	productsList := c.productService.List()
	for _, v := range productsList {
		outputMsg += v.Title + "\n\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).ListCommand
}
