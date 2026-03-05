package commands

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) List(inputMessage *tgbotapi.Message) {
	products := commander.productService.List()

	msgText := ""

	for _, p := range products {
		msgText += p.Title
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	msg.ReplyToMessageID = inputMessage.MessageID

	serialisedData, err := json.Marshal(Command{
		T:     "list",
		From:  5,
		Count: 5,
	})

	if err != nil {
		log.Println("marshaling error")
		commander.serverErrorHandler(inputMessage)
		return
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serialisedData)),
		),
	)

	_, err = commander.bot.Send(msg)

	if err != nil {
		log.Println("Send error")
	}
}
