package telegrambot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/yahahaff/rapide/pkg/logger"
)

type Bot struct {
	api    *tgbotapi.BotAPI
	config tgbotapi.UpdateConfig
}

var Bots *Bot

func NewBot(token string) error {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot := &Bot{
		api:    api,
		config: tgbotapi.NewUpdate(0),
	}

	bot.api.RemoveWebhook()
	bot.api.GetMe()

	Bots = bot

	return nil
}

func (bot *Bot) Start() {

	go func() {
		updates, err := bot.api.GetUpdatesChan(bot.config)
		if err != nil {
			logger.ErrorString("", "Error", err.Error())
			return
		}
		logger.DebugString("TelegramBot", "message", "TelegramBot started successfully")
		for update := range updates {
			if update.Message == nil {
				continue
			}

			bot.HandleMessage(update.Message)
		}
	}()
}

func (bot *Bot) HandleMessage(message *tgbotapi.Message) {
	logger.DebugString("TelegramBot", "message", fmt.Sprintf("Received message from chat ID %d: %s", message.Chat.ID, message.Text))
	if message.IsCommand() {
		switch message.Command() {
		case "harley":
			bot.SendMessage(message.Chat.ID, "👅🐕 👅🐕 👅🐕 👅🐕 👅🐕")
		case "link":
			bot.SendMessage(message.Chat.ID, "🐔 🐔 🐔 🐔 🐔 🐔")
		case "shine":
			bot.SendMessage(message.Chat.ID, "🐺 🐺 🐺 🐺 🐺 🐺")
		case "help":
			bot.SendMessage(message.Chat.ID, "/harley /link /help")
		default:
			bot.SendMessage(message.Chat.ID, "Unknown command")
		}
	}
}

func (bot *Bot) SendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.api.Send(msg)
	return err
}
