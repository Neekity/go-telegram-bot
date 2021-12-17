package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/syyongx/php2go"
	"log"
	"neekity.com/go-telegram-bot/src/internal"
	configType "neekity.com/go-telegram-bot/src/internal/config"
	"neekity.com/go-telegram-bot/src/internal/plugins"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(configType.Conf.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			command := update.Message.Text[0:2]
			if !php2go.InArray(update.Message.Chat.ID, configType.Conf.SupportChatIds) {
				command = "error"
			}
			switch command {
			case "bb":
				internal.SendPhotos(bot, update.Message.Chat.ID, plugins.GetRandomPic("bb", update.Message.Text[2:]))
			case "bt":
				internal.SendPhotos(bot, update.Message.Chat.ID, plugins.GetRandomPic2("bt", update.Message.Text[2:]))
			case "as":
				internal.SendResources(bot, update.Message.Chat.ID, plugins.GetRandomResource("as", update.Message.Text[2:]))

			case "error":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "errors")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
