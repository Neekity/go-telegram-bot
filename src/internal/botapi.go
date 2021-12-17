package internal

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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
	bot.SendMediaGroup(mediaGroup)
}

func sendPhoto(bot *tgbotapi.BotAPI, chatId int64, picUrl string) {
	msg := tgbotapi.NewPhoto(chatId, tgbotapi.FileURL(picUrl))
	bot.Send(msg)
}
