package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"neekity.com/go-telegram-bot/src/internal/plugins"
)

func SendPhotos(bot *tgbotapi.BotAPI, chatId int64, picUrls []string) {
	if len(picUrls) > 10 {
		for _, url := range picUrls {
			sendPhoto(bot, chatId, url)
		}
		return
	}
	var medias []interface{}
	for _, url := range picUrls {
		file := tgbotapi.FileURL(url)
		medias = append(medias, tgbotapi.NewInputMediaPhoto(file))
	}

	mediaGroup := tgbotapi.NewMediaGroup(chatId, medias)
	sendMediaGroup(bot, chatId, mediaGroup)

}

func SendResources(bot *tgbotapi.BotAPI, chatId int64, resourceUrls []plugins.RandomResource) {
	var medias []interface{}
	for _, url := range resourceUrls {
		file := tgbotapi.FileURL(url.FileUrl)
		if len(file) == 0 {
			continue
		}
		if url.FileExt == "mp4" {
			msg := tgbotapi.NewVideo(chatId, file)

			sendMessage(bot, chatId, msg)
		} else if len(medias) == 10 {
			msg := tgbotapi.NewPhoto(chatId, file)

			sendMessage(bot, chatId, msg)
		} else {
			medias = append(medias, tgbotapi.NewInputMediaPhoto(file))
		}
	}

	mediaGroup := tgbotapi.NewMediaGroup(chatId, medias)
	sendMediaGroup(bot, chatId, mediaGroup)
}

func sendPhoto(bot *tgbotapi.BotAPI, chatId int64, picUrl string) {
	msg := tgbotapi.NewPhoto(chatId, tgbotapi.FileURL(picUrl))

	sendMessage(bot, chatId, msg)

}

func SendError(bot *tgbotapi.BotAPI, chatId int64, err error) {
	msg := tgbotapi.NewMessage(chatId, err.Error())
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
		bot.Send(msg)
	}
}

func sendMediaGroup(bot *tgbotapi.BotAPI, chatId int64, mediaGroup tgbotapi.MediaGroupConfig) {
	if _, err := bot.SendMediaGroup(mediaGroup); err != nil {
		SendError(bot, chatId, err)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatId int64, msg tgbotapi.Chattable) {
	if _, err := bot.Send(msg); err != nil {
		SendError(bot, chatId, err)
	}
}
