package comands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Add(inputMessage *tgbotapi.Message) {

	c.productService.Add(inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Element added",
	)

	c.bot.Send(msg)
}
