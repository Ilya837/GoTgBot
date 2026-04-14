package product

import (
	"log"
	"strconv"
	"strings"

	"github.com/Ilya837/GoTgMod/internal/model/market"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) New(inputMessage *tgbotapi.Message) {
	title := inputMessage.CommandArguments()

	if title == "" || strings.Count(title, " ") != 0 {
		commander.WrongFormat(inputMessage, "/new__market__product title")
		return
	}

	id, err := (*commander.Service).Create(market.Product{Title: title})

	if err != nil {
		log.Println("error at create: " + err.Error())
		commander.ServerError(inputMessage)
		return
	}

	commander.sendMessage("product was created with id "+strconv.Itoa(int(id)), inputMessage)

}
