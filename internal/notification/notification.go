package notification

import "github.com/nikoksr/notify/service/telegram"

func InitReceivers(tgBotToken string, tgChatID int64) *telegram.Telegram {

	return initTelegram(tgBotToken, tgChatID)

}

func initTelegram(tgBotToken string, tgChatID int64) *telegram.Telegram {
	telegramService, _ := telegram.New(tgBotToken)
	telegramService.AddReceivers(tgChatID)

	return telegramService
}
