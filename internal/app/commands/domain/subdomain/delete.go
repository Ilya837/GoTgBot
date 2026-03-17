package subdomain

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	if args == "" {
		commander.WrongFormat(inputMessage, "/delete__domain__subdomain id")
		return
	}

	id, err := strconv.Atoi(args)

	if err != nil {
		commander.sendMessage("Id must be number", inputMessage)
		return
	}

	if id < 0 {
		commander.sendMessage("Id must be more then 0", inputMessage)
		return
	}

	err = (*commander.Service).Remove(uint64(id))

	if err != nil {
		log.Println("error in Remove method")
		commander.ServerError(inputMessage)
		return
	}

	commander.sendMessage("product deleted", inputMessage)
}
