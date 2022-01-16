package telegram_test

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicourrrn/fin-manager/ui/telegram"
	"log"
	"os"
	"testing"
)

func TestConn(t *testing.T) {
	bot, err := telegram.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		t.Error(err)
	}
	bot.AddCommand("/start", func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
		newMsg := tgbotapi.NewMessage(msg.From.ID, "Hello from bot")
		_, err := bot.Send(newMsg)
		return err
	})
	errChan := make(chan error)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.Bot.GetUpdatesChan(u)
	go bot.Handle(updates, errChan)
	for err = range errChan {
		if err != nil {
			log.Println(err)
		}
	}
}
