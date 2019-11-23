package main

import (
	"net/http"

	"komodorIntegrations/config"
	"komodorIntegrations/db"
	_ "net/http/pprof"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"addr": config.Addr,
	}).Info("Komodor Integrations starting")
	_, err := db.NewDbClient()
	if err != nil {
		log.WithError(err).Fatal("Failed to init db")
	}
	s := InitHttpServer()
	if err := http.ListenAndServe(config.Addr, s.Mux); err != nil {
		log.WithError(err).WithField("addr", config.Addr).Fatal("Listen and serve")
	}

}
