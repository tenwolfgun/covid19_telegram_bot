package handlers

import (
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Bot **
type Bot struct {
	*tb.Bot
}

// NewBot **
func NewBot() (*Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 60 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	return &Bot{b}, nil
}
