package product

import (
	"log"
	"strconv"
	"strings"

	"github.com/Ilya837/GoTgMod/internal/model/market"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	split := strings.Split(args, " ")

	if len(split) > 2 {
		commander.WrongFormat(inputMessage, "/edit__market__product id title")
		return
	}

	id, err := strconv.Atoi(split[0])

	if err != nil {
		commander.sendMessage("Id must be number: ", inputMessage)
		return
	}

	if id < 0 {
		commander.sendMessage("Id must be more then 0", inputMessage)
		return
	}

	err = (*commander.Service).Update(uint64(id), market.Product{Title: split[1]})

	if err != nil {
		log.Println("error in Update method")
		commander.ServerError(inputMessage)
		return
	}

	commander.sendMessage("product updated", inputMessage)

}
