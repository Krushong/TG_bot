package comands

import (
	"TGbot/bot/internal/product"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil || update.CallbackQuery != nil { // If we got a message

		//Вариант JSON

		if update.CallbackQuery != nil {
			parsedData := CommandData{}
			json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
			msg := tgbotapi.NewMessage(
				update.CallbackQuery.Message.Chat.ID,
				fmt.Sprintf("Parsed: %s\n", parsedData),
			)
			c.bot.Send(msg)
			return
		}

		//Вариант не джейсон

		//if update.CallbackQuery != nil {
		//	args := strings.Split(update.CallbackQuery.Data, "_")
		//	msg := tgbotapi.NewMessage(
		//		update.CallbackQuery.Message.Chat.ID,
		//		fmt.Sprintf("Command: %s\n", args[0])+
		//			fmt.Sprintf("Offset: %s\n", args[1]),
		//	)
		//	c.bot.Send(msg)
		//	return
		//}

		if update.Message == nil {
			return
		}

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		case "delete":
			c.Delete(update.Message)
		case "add":
			c.Add(update.Message)
		case "edit":
			c.Edit(update.Message)
		case "xo":
			c.xo(update.Message)
		default:
			if product.EditElement == true {
				c.edit_value(update.Message)
			} else if product.Bollgame == true {
				c.move(update.Message)
			} else {
				c.Default(update.Message)
			}
		}
	}
}

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}
