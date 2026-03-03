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
		log.Println("wrong args", args)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprint("wrong args", args))
		msg.ReplyToMessageID = inputMessage.MessageID
		commander.bot.Send(msg)
		return
	}

	product, err := commander.productService.Get(idx)

	if err != nil {
		msgText := fmt.Sprintf("fail to get product with idf: %d: %v", idx, err)
		log.Printf(msgText)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
		msg.ReplyToMessageID = inputMessage.MessageID
		commander.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf(product.Title))

	msg.ReplyToMessageID = inputMessage.MessageID
	commander.bot.Send(msg)

}
