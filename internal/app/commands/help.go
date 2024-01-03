package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
			"/list - list of items\n"+
			"/get - get list of arguments",
	)

	c.bot.Send(msg)
}
