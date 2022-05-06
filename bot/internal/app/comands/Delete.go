package comands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	c.productService.Delete(idx)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Element delete",
	)

	c.bot.Send(msg)
}
