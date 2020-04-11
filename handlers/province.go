package handlers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/w0rm1995/covid19_telegram_bot/models"
	tb "gopkg.in/tucnak/telebot.v2"
)

// HandleProvince **
func (b *Bot) HandleProvince(m *tb.Message) {
	province := models.Province{}

	resp, err := province.GetCovid19ByProvince()

	if err != nil {
		log.Infoln(err)
	}

	if !m.Private() {
		return
	}

	for i := 0; i < len(resp.Data); i++ {
		b.Send(m.Sender, fmt.Sprintf(`Provinsi: %v
	
Terinfkesi: %v
Sembuh: %v
Meninggal: %v`, resp.Data[i].Provinsi, resp.Data[i].KasusPosi, resp.Data[i].KasusSemb, resp.Data[i].KasusMeni))
		log.Infoln(m.Payload)
	}
}
