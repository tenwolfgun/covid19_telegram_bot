package handlers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/w0rm1995/covid19_telegram_bot/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

// HandleDaily **
func (b *Bot) HandleDaily(m *tb.Message) {
	daily := models.Daily{}

	resp, err := daily.GetDailyCovid19()

	if err != nil {
		log.Infoln(err)
	}

	if !m.Private() {
		return
	}

	for i := 0; i < len(resp)-49; i++ {
		b.Send(m.Sender, fmt.Sprintf(`Tanggal: %v
				
Terinfeksi:  %v
Sembuh:  %v`, resp[i].ReportDateString, resp[i].TotalConfirmed, resp[i].TotalRecovered))
	}

	log.Infoln(m.Payload)
}
