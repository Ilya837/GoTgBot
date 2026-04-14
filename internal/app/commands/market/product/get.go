package product

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (commander Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		commander.WrongFormat(inputMessage, "/get__market__product id")
		return
	}

	if idx < 0 {
		log.Printf("id must be more than 0: %s", args)
		commander.sendMessage(
			fmt.Sprintf("id must be more than 0: %s", args),
			inputMessage)
		return
	}

	product, err := (*commander.Service).Describe(uint64(idx))

	if err != nil {
		commander.sendMessage(
			fmt.Sprintf("fail to get product with id %d: %v", idx, err),
			inputMessage)
		return
	}

	if product == nil {
		commander.sendMessage(
			fmt.Sprintf("product with id %d is not exist", idx),
			inputMessage)
		return
	}

	commander.sendMessage(
		product.Title,
		inputMessage)

}
