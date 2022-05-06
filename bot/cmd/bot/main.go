package main

import (
	"TGbot/bot/internal/app/comands"
	"TGbot/bot/internal/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("5376160580:AAHAC9GQnLhxWiQjK1OAJteVQuf0lQ7aBaM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	commander := comands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
