package handlers

import (
	log "github.com/sirupsen/logrus"

	tb "gopkg.in/tucnak/telebot.v2"
)

// HandleSource **
func (b *Bot) HandleSource(m *tb.Message) {

	if !m.Private() {
		return
	}

	b.Send(m.Sender, `Sumber data: 
	
Global: https://covid19.mathdro.id/api/
Indonesia : https://indonesia-covid-19.mathdro.id/api/
KalSel : http://corona.kalselprov.go.id/cov_map`, tb.NoPreview)

	log.Infoln(m.Payload)
}
