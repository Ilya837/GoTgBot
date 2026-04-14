package product

import (
	service "github.com/Ilya837/GoTgMod/internal/service/market/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProductCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	Default(inputMsg *tgbotapi.Message)
	CallBack(inputMsg *tgbotapi.CallbackQuery)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
	ServerError(inputMessage *tgbotapi.Message)
	WrongFormat(inputMessage *tgbotapi.Message, rightFormat string)
}

type Commander struct {
	bot     *tgbotapi.BotAPI
	Service *service.ProductService
}

func NewProductCommander(bot *tgbotapi.BotAPI, service service.ProductService) ProductCommander {
	return &Commander{bot: bot, Service: &service}
}
