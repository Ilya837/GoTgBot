package commands

import (
	"encoding/json"
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
	T     string `json:"type"`
	From  int    `json:"from"`
	Count int    `json:"count"`
}

func (commander Commander) HandleUpdate(update *tgbotapi.Update) {

	defer commander.panicHandler(update)

	if update.CallbackQuery != nil {

		log.Printf("[%s] %s", update.CallbackQuery.Message.From.UserName, update.CallbackQuery.Data)

		command := Command{}

		err := json.Unmarshal([]byte(update.CallbackQuery.Data), &command)

		if err != nil {
			log.Println("unmarshaling error")
			commander.serverErrorHandler(update.CallbackQuery.Message)
			return
		}

		commander.sendMessage(
			fmt.Sprintf("Data: %+v", command),
			update.CallbackQuery.Message)

		return
	}

	if update.Message != nil {

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

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

		return
	}

}

func (commander Commander) panicHandler(update *tgbotapi.Update) {
	if panicValue := recover(); panicValue != nil {
		log.Println("recovered from panic: ", panicValue)

		if update.CallbackQuery != nil {

			commander.serverErrorHandler(update.CallbackQuery.Message)

		} else if update.Message != nil {

			commander.serverErrorHandler(update.Message)
		}
	}
}

func (commander Commander) serverErrorHandler(inputMessage *tgbotapi.Message) {

	commander.sendMessage(
		"server error",
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
