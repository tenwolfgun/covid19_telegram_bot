package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/w0rm1995/covid19_telegram_bot/handlers"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	b, err := handlers.NewBot()

	if err != nil {
		log.Infoln(err)
	}

	b.Handle("/start", b.HandleStart)

	b.Handle("/global", b.HandleGlobal)

	b.Handle("/harian", b.HandleDaily)

	b.Handle("/indonesia", b.HandleIndonesia)

	b.Handle("/provinsi", b.HandleProvince)

	b.Handle("/sumber", b.HandleSource)

	b.Handle("/kalsel", b.HandleKalSel)

	b.Handle("/telepon", b.HandleCallCenter)

	log.Infoln("Running")
	b.Start()
}
