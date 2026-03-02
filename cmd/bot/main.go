package main

import (
	"log"
	"os"
	"strings"

	"github.com/Ilya837/GoTgMod/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message != nil { // If we got a message

			switch update.Message.Command() {
			case "help":
				Help(bot, update.Message)
			case "list":
				ListHandler(bot, update.Message, *productService)
			default:
				defaultHandler(bot, update.Message)
			}

		}
	}
}

func ListHandler(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService product.Service) {
	products := productService.List()

	msgText := ""

	for _, p := range products {
		msgText += p.Title
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}

func defaultHandler(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)

}

func Help(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		strings.Join([]string{
			"/help - help",
			"/list - list"},
			"\n"))
	msg.ReplyToMessageID = inputMessage.MessageID

	bot.Send(msg)
}
