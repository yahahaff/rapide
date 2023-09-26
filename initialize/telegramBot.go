package initialize

import (
	"github.com/yahahaff/rapide/pkg/config"
	"github.com/yahahaff/rapide/pkg/telegrambot"
)

// SetupTelegramBot 初始化TelegramBot
func SetupTelegramBot() {
	// 初始化TelegramBot
	err := telegrambot.NewBot(config.GetString("telegram.bot_token"))
	if err != nil {
		return
	}
	telegrambot.Bots.Start()
	// go telegrambot.Bots.Start()
}
