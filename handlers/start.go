package handlers

import (
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

// HandleStart **
func (b *Bot) HandleStart(m *tb.Message) {
	if !m.Private() {
		return
	}

	b.Send(m.Sender, `/global => - Informasi global kasus covid 19

/harian => - Informasi harian kasus covid 19

/indonesia => - Informasi kasus covid 19 di Indonesia
	
/provinsi => - Informasi kasus covid 19 berdasarkan provinsi di Indonesia
	
/sumber => - Informasi mengenai sumber data
	
/kalsel => - Informasi kasus covid 19 di Kalimantan Selatan
	
/telepon => - Informasi Call Center Covid 19 di Kalimantan Selatan`)
	log.Infoln(m.Payload)
}
