package handlers

import (
	"encoding/json"
	"komodorIntegrations/db"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Request struct {
	Message string `json:"message"`
	Author  string `json:"author"`
}

func incomingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var d Request
		err := decoder.Decode(&d)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = handleIncoming(d)
		if err != nil {
			log.WithError(err).Warn("load handler response")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleIncoming(req Request) error {
	details, err := db.GetUserDetails(req.Author)
	if err != nil {
		return err
	}
	if details == nil {
		return errors.Newf("Coudln't find user: %s", req.Author)
	}
	_ = SendSlackMessage(req.Message, details)
	return nil
}
