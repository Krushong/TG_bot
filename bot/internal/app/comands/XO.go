package comands

import (
	"TGbot/bot/internal/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var letters []string
var board string

func CreateCleanBoard(board string) string {
	letters = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}

	board = "0|1|2\n" +
		"_____\n" +
		"3|4|5\n" +
		"_____\n" +
		"6|7|8\n"
	return board
}

func (c *Commander) xo(inputMessage *tgbotapi.Message) {

	//args := inputMessage.CommandArguments()
	//
	//idx, err := strconv.Atoi(args)
	//if err != nil {
	//	log.Println("wrong args", args)
	//	return
	//}
	//
	//c.productService.Xo(idx)

	product.Bollgame = true

	board := CreateCleanBoard("")

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		board+"\nВведите число где будет (О)",
	)

	c.bot.Send(msg)
}
