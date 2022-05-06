package comands

import (
	"TGbot/bot/internal/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) edit_value(inputMessage *tgbotapi.Message) {

	c.productService.Edit_value(inputMessage)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Value edited",
	)

	product.EditElement = false
	c.bot.Send(msg)
}
