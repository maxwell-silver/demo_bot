package main

import (
	"demo_bot/internal/service/product"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "/help - help")

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func listCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message, productService *product.Service) {
	outputMsg := "There are all products:\n\n"
	products := productService.List()

	for _, p := range products {
		outputMsg += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsg)

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("You wrote: %s", inputMsg.Text))

	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
