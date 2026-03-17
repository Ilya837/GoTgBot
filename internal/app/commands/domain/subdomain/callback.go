package subdomain

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) CallBack(callbackQuery *tgbotapi.CallbackQuery) {

	limit := 5

	log.Printf("[%s] %s", callbackQuery.Message.From.UserName, callbackQuery.Data)

	command := Command{}

	err := json.Unmarshal([]byte(callbackQuery.Data), &command)

	if err != nil {
		log.Println("unmarshaling error")
		commander.ServerError(callbackQuery.Message)
		return
	}

	list, err := (*commander.Service).List(uint64(command.Cursor), uint64(command.Limit))

	if err != nil {
		//
	}

	msgText := ""

	for _, p := range list {
		msgText += p.Title
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, msgText)

	if len(list) == limit {

		list2, err := (*commander.Service).List(uint64(command.Cursor), uint64(command.Limit+1))

		if err != nil {
			//
		}

		if len(list2) != len(list) {

			serialisedData, err := json.Marshal(Command{
				T:      "list",
				Cursor: command.Cursor + command.Limit,
				Limit:  limit,
			})

			if err != nil {
				log.Println("marshaling error")
				commander.ServerError(callbackQuery.Message)
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
