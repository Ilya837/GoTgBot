package subdomain

import (
	service "github.com/Ilya837/GoTgMod/internal/service/domain/subdomain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubdomainCommander interface {
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
	Service *service.SubdomainService
}

func NewSubdomainCommander(bot *tgbotapi.BotAPI, service service.SubdomainService) SubdomainCommander {
	return &Commander{bot: bot, Service: &service}
}
