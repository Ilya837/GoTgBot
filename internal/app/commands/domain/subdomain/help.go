package subdomain

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Help(inputMessage *tgbotapi.Message) {

	commander.sendMessage(
		strings.Join([]string{
			"/help__domain__subdomain — print list of commands",
			"/get__domain__subdomain — get a entity",
			"/list__domain__subdomain — get a list of your entity",
			"/delete__domain__subdomain — delete an existing entity",
			"/new__domain__subdomain — create a new entity",
			"/edit__domain__subdomain — edit a entity"},
			"\n"),
		inputMessage)
}
