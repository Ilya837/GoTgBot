package main

import (
	"log"
	"os"

	commands "github.com/Ilya837/GoTgMod/internal/app/commands/market/product"
	service "github.com/Ilya837/GoTgMod/internal/service/market/product"
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

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := service.NewService()

	commander := commands.NewProductCommander(bot, productService)

	for update := range updates {

		handleUpdate(&commander, &update)

	}
}

func handleUpdate(commander *commands.ProductCommander, update *tgbotapi.Update) {

	defer panicHandler(commander, update)

	if update.CallbackQuery != nil {

		(*commander).CallBack(update.CallbackQuery)
	}

	if update.Message != nil {

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "help__market__product":
			(*commander).Help(update.Message)
		case "list__market__product":
			(*commander).List(update.Message)
		case "get__market__product":
			(*commander).Get(update.Message)
		case "delete__market__product":
			(*commander).Delete(update.Message)
		case "new__market__product":
			(*commander).New(update.Message)
		case "edit__market__product":
			(*commander).Edit(update.Message)
		default:
			(*commander).Default(update.Message)
		}

		return
	}

}

func panicHandler(commander *commands.ProductCommander, update *tgbotapi.Update) {
	if panicValue := recover(); panicValue != nil {
		log.Println("recovered from panic: ", panicValue)

		if update.CallbackQuery != nil {

			(*commander).ServerError(update.CallbackQuery.Message)

		} else if update.Message != nil {

			(*commander).ServerError(update.Message)
		}
	}
}
