package product

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Help(inputMessage *tgbotapi.Message) {

	commander.sendMessage(
		strings.Join([]string{
			"/help__market__product — print list of commands",
			"/get__market__product — get a entity",
			"/list__market__product — get a list of your entity",
			"/delete__market__product — delete an existing entity",
			"/new__market__product — create a new entity",
			"/edit__market__product — edit a entity"},
			"\n"),
		inputMessage)
}
