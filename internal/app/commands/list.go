package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	outputMsg := "There are all products:\n\n"
	products := c.productService.List()

	for _, p := range products {
		outputMsg += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsg)

	c.bot.Send(msg)
}
