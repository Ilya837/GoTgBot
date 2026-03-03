package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID

	commander.bot.Send(msg)

}

func (commander Commander) HandleUpdate(update *tgbotapi.Update) {

	defer commander.panicHandler(update.Message)

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		commander.Help(update.Message)
	case "list":
		commander.List(update.Message)
	case "get":
		commander.Get(update.Message)
	default:
		commander.Default(update.Message)
	}

}

func (commander Commander) panicHandler(inputMessage *tgbotapi.Message) {
	if panicValue := recover(); panicValue != nil {
		log.Println("recovered from panic: ", panicValue)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "server error")
		msg.ReplyToMessageID = inputMessage.MessageID

		commander.bot.Send(msg)
	}
}
