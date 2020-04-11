package handlers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/w0rm1995/covid19_telegram_bot/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

// HandleIndonesia **
func (b *Bot) HandleIndonesia(m *tb.Message) {
	country := models.Indonesia{}

	resp, err := country.GetIndonesiaCovid19()

	if err != nil {
		log.Infoln(err)
	}

	if !m.Private() {
		return
	}

	b.Send(m.Sender, fmt.Sprintf(`TERINFEKSI:  %v, SEMBUH:  %v, PERAWATAN: %v, MENINGGAL:  %v`, resp.JumlahKasus, resp.Sembuh, resp.Perawatan, resp.Meninggal))
	log.Infoln(m.Payload)
}
