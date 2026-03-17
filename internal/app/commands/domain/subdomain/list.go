package subdomain

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) List(inputMessage *tgbotapi.Message) {

	limit := 5

	list, err := (*commander.Service).List(0, uint64(limit))

	if err != nil {
		//
	}

	msgText := ""

	for _, p := range list {
		msgText += p.Title
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	msg.ReplyToMessageID = inputMessage.MessageID

	if len(list) == limit {

		list2, err := (*commander.Service).List(0, uint64(limit+1))

		if err != nil {
			//
		}

		if len(list) != len(list2) {

			serialisedData, err := json.Marshal(Command{
				T:      "list",
				Cursor: 5,
				Limit:  5,
			})

			if err != nil {
				log.Println("marshaling error")
				commander.ServerError(inputMessage)
				return
			}

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Next page", string(serialisedData)),
				),
			)
		}
	}

	_, err = commander.bot.Send(msg)

	if err != nil {
		log.Println("Send error")
	}
}
