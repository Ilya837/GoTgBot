package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (commander Commander) List(inputMessage *tgbotapi.Message) {
	products := commander.productService.List()

	msgText := ""

	for _, p := range products {
		msgText += p.Title
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	msg.ReplyToMessageID = inputMessage.MessageID

	commander.bot.Send(msg)
}
