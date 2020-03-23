package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/w0rm1995/covid19_telegram_bot/models"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"time"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 60 * time.Second},
	})

	if err != nil {
		log.Infoln(err)
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		b.Send(m.Sender, `/global => - Informasi global kasus covid 19

	/harian => - Informasi harian kasus covid 19

	/indonesia => - Informasi kasus covid 19 di Indonesia
	
	/provinsi => - Informasi kasus covid 19 berdasarkan provinsi di Indonesia
	
	/sumber => - Informasi mengenai sumber data`)
		log.Infoln(m.Payload)
	})

	b.Handle("/global", func(m *tb.Message) {
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
	})

	b.Handle("/harian", func(m *tb.Message) {
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
	})

	b.Handle("/indonesia", func(m *tb.Message) {
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
	})

	b.Handle("/image", func(m *tb.Message) {
		image := &tb.Photo{
			File:      tb.FromURL("https://covid19.mathdro.id/api/og"),
			ParseMode: tb.ParseMode(tb.ModeMarkdown),
			Caption:   "https://covid19.mathdro.id/api/og",
		}

		b.Send(m.Sender, image)

		log.Infoln(m.Payload)
	})

	b.Handle("/provinsi", func(m *tb.Message) {
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
	})

	b.Handle("/sumber", func(m *tb.Message) {

		if err != nil {
			log.Infoln(err)
		}

		if !m.Private() {
			return
		}

		b.Send(m.Sender, `Sumber data: 
	
	Global: https://covid19.mathdro.id/api/
	Indonesia : https://indonesia-covid-19.mathdro.id/api/`, tb.NoPreview)

		log.Infoln(m.Payload)
	})

	log.Infoln("Running")
	b.Start()
}
