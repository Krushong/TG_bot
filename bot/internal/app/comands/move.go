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
		board = strings.Replace(board, OldChar, NewChar, 1)
		fmt.Println(board)

		if boardBefor != board {
			mapStep[OldChar] = NewChar
			if proverka("X") {
				return (board + "\nТы выиграл красава игра закончина введи заново /xo")
			} else {
				return StepComp("O")
			}
		}
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
	fmt.Println(NumberGameChar)
	if NumberGameChar < 0 || NumberGameChar > 9 {
		errtext := fmt.Sprintf("Вы ввели %s Введите число от 1 до 9 или /EndGame", GameChar)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			errtext,
		)

		c.bot.Send(msg)
		return
	}

	msgboard := Step("X", GameChar)
	fmt.Println(msgboard)
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgboard,
	)

	c.bot.Send(msg)
}
