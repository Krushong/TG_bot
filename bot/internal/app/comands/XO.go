package comands

import (
	"TGbot/bot/internal/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var mapStep map[string]string
var board string

func CreateCleanBoard(board string) string {
	mapStep = map[string]string{"0": "-", "1": "-", "3": "-", "4": "-", "5": "-", "6": "-", "7": "-", "8": "-"}

	board = "0|1|2\n" +
		"_____\n" +
		"3|4|5\n" +
		"_____\n" +
		"6|7|8\n"
	return board
}

func (c *Commander) xo(inputMessage *tgbotapi.Message) {

	product.Bollgame = true

	board = CreateCleanBoard("")

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		board+"\nВведите число где будет (X)",
	)

	c.bot.Send(msg)
}
