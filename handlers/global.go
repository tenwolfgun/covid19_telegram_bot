package handlers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/w0rm1995/covid19_telegram_bot/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

// HandleGlobal **
func (b *Bot) HandleGlobal(m *tb.Message) {
	current := models.CurrentCase{}

	resp, err := current.GetCurrentCovid19()

	if err != nil {
		log.Infoln(err)
	}

	if !m.Private() {
		return
	}

	b.Send(m.Sender, fmt.Sprintf(`TERINFEKSI:  %v, SEMBUH:  %v, MENINGGAL:  %v`, resp.Confirmed.Value, resp.Recovered.Value, resp.Deaths.Value))
	log.Infoln(m.Payload)
}
