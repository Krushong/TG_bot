package comands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func StepComp(NewChar, OldChar string) string {

}

func Step(NewChar, OldChar string) string {
	boardBefor := board
	strings.Replace(board, OldChar, NewChar, 1)
	if boardBefor == board {

		return StepComp()
	} else {
		return "Введен повторяющийся символ"
	}
}

func (c *Commander) move(inputMessage *tgbotapi.Message) {

	GameChar := inputMessage.Text
	if GameChar != "X" || GameChar != "O" {

	}
	//Step()

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Element delete",
	)

	c.bot.Send(msg)
}
