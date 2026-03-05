package commands

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
		log.Printf("wrong args: %s", args)

		commander.sendMessage(
			fmt.Sprintf("wrong args: %s", args),
			inputMessage)
		return
	}

	product, err := commander.productService.Get(idx)

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
