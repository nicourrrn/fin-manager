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

func NewBot(token string) (*TgBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	me, err := bot.GetMe()
	if err != nil {
		return nil, err
	}
	log.Printf("%s loggined", me.FirstName)
	return &TgBot{
		Bot:      bot,
		Commands: make(map[string]BotHandler),
		Statuses: make(map[int64]BotHandler),
	}, nil
}

func (b *TgBot) AddCommand(command string, handler BotHandler) {
	b.Commands[command] = handler
}

func (b *TgBot) Handle(upds tgbotapi.UpdatesChannel, errChan chan error) {
	var (
		//upd *tgbotapi.Update
		msg *tgbotapi.Message
	)

	for upd := range upds {
		if upd.Message == nil {
			continue
		}
		msg = upd.Message
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
