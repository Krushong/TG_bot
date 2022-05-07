package comands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help \n/list - list"+
			"\n/get - get \n/add - add \n/xo - XO")

	c.bot.Send(msg)
}

func init() {
	registeredComands["help"] = (*Commander).Help
}
