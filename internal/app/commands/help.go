package commands

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		strings.Join([]string{
			"/help - help",
			"/list - list products",
			"/get id - get product by id"},
			"\n"))
	msg.ReplyToMessageID = inputMessage.MessageID

	commander.bot.Send(msg)
}
