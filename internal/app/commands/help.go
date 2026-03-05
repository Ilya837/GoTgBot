package commands

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Help(inputMessage *tgbotapi.Message) {

	commander.sendMessage(
		strings.Join([]string{
			"/help - help",
			"/list - list products",
			"/get id - get product by id"},
			"\n"),
		inputMessage)
}
