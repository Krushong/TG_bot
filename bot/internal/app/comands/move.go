package comands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func StepComp(NewChar string) string {
	booltest := true
	var OldChar string
	for booltest {
		OldChar = strconv.Itoa(rand.Intn(9))
		if mapStep[OldChar] != "X" && mapStep[OldChar] != "O" {
			booltest = false
		}
	}

	strings.Replace(board, OldChar, NewChar, 1)
	mapStep[OldChar] = NewChar
	return board
}

func Step(NewChar, OldChar string) string {
	boardBefor := board
	if mapStep[OldChar] != "X" && mapStep[OldChar] != "O" {
		strings.Replace(board, OldChar, NewChar, 1)
	}
	if boardBefor == board {
		mapStep[OldChar] = NewChar
		return StepComp("O")
	} else {
		return "Введен повторяющийся символ"
	}
}

func (c *Commander) move(inputMessage *tgbotapi.Message) {

	GameChar := inputMessage.Text
	NumberGameChar, err := strconv.Atoi(GameChar)
	if err != nil {
		log.Fatal(err)
	}

	if NumberGameChar < 1 || NumberGameChar > 9 {
		errtext := fmt.Sprintf("Вы ввели %s Введите число от 1 до 9 или /EndGame", GameChar)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			errtext,
		)

		c.bot.Send(msg)
		return
	}

	Step("X", GameChar)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		board,
	)

	c.bot.Send(msg)
}
