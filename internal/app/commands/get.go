package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v\n", idx, err)
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		product.Title,
	)

	c.bot.Send(msg)
}
