package subdomain

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID

	commander.bot.Send(msg)

}

type Command struct {
	T      string `json:"type"`
	Cursor int    `json:"cursor"`
	Limit  int    `json:"limit"`
}

func (commander Commander) ServerError(inputMessage *tgbotapi.Message) {

	commander.sendMessage(
		"server error",
		inputMessage)
}

func (commander Commander) WrongFormat(inputMessage *tgbotapi.Message, rightFormat string) {

	commander.sendMessage(
		"Wrong format. Right format is: "+rightFormat,
		inputMessage)
}

func (commander Commander) sendMessage(text string, inputMessage *tgbotapi.Message) {

	msgText := fmt.Sprintln(text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	msg.ReplyToMessageID = inputMessage.MessageID

	_, err := commander.bot.Send(msg)

	if err != nil {
		log.Println("Send error")
		return
	}

	log.Printf("request sended")

}
