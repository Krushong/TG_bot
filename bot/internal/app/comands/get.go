package comands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Println("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	c.bot.Send(msg)
}

func init() {
	registeredComands["get"] = (*Commander).Get
}
