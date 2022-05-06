package comands

import (
	"TGbot/bot/internal/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	c.productService.Edit(idx)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Please enter text",
	)

	product.EditElement = true
	c.bot.Send(msg)
}
