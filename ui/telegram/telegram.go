package telegram

import (
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type BotHandler func(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error

type TgBot struct {
	Bot      *tgbotapi.BotAPI
	Commands map[string]BotHandler
	Statuses map[int64]BotHandler
}

func NewBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	me, err := bot.GetMe()
	if err != nil {
		return nil, err
	}
	log.Printf("%s loggined", me.FirstName)
	return bot, nil
}

func (b *TgBot) AddCommand(command string, handler BotHandler) {
	b.Commands[command] = handler
}

func (b *TgBot) Handle(msgs chan *tgbotapi.Message, errChan chan error) {
	var msg *tgbotapi.Message
	for msg = range msgs {
		if h, ok := b.Commands[msg.Text]; ok {
			err := h(b.Bot, msg)
			if err != nil {
				errChan <- err
			}
		} else if h, ok := b.Statuses[msg.From.ID]; ok {
			err := h(b.Bot, msg)
			if err != nil {
				errChan <- err
			}
		} else {
			errChan <- errors.New("command and status not found")
		}
	}

}
