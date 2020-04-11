package handlers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/w0rm1995/covid19_telegram_bot/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

// HandleKalSel **
func (b *Bot) HandleKalSel(m *tb.Message) {
	kalsel := models.KalSel{}

	resp, err := kalsel.GetKalSelCovid19()

	if err != nil {
		log.Infoln(err)
	}

	if !m.Private() {
		return
	}

	for i := 0; i < len(resp); i++ {
		b.Send(m.Sender, fmt.Sprintf(`Kabupaten/Kota: %v
	
Orang Dalam Pemantauan: %v
Pasien Dalam Pengawasan: %v
Positif: %v
Sembuh: %v
Meninggal: %v`, resp[i].Name, resp[i].CovOdpCount, resp[i].CovPdpCount, resp[i].CovPositiveCount, resp[i].CovRecoveredCount, resp[i].CovDiedCount))
	}

	log.Infoln(m.Payload)
}
