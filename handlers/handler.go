package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"komodorIntegrations/db"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var funcHandlersMap = map[string]func(*Request) ([]byte, error){
	"slack":           SetSlackMessage,
	"generic_webhook": SetRestMessage,
	"pagerduty":       SetPagerdutyMessage,
}

type Request struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Details     string `json:"details"`
	Username    string `json:"username"`
	// Service ID Is Currently being used only for PD
	ServiceId string `json:"service_id"`
}

func IncomingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var d Request
		err := decoder.Decode(&d)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = handleIncoming(&d)
		if err != nil {
			log.WithError(err).Warn("load handler response")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleIncoming(req *Request) error {
	webhooks, err := db.Client.GetWebhooks(req.Username)
	if err != nil {
		return err
	}
	if webhooks == nil {
		return errors.New(fmt.Sprintf("Coudln't find user: %s", req.Username))
	}
	for app, appDetails := range webhooks.Apps {
		handler := funcHandlersMap[app]
		message, err := handler(req)
		log.Info(string(message))
		if err != nil {
			log.WithError(err).Error("Failed to call handler")
			continue
		}

		if err = sendRequest(message, &appDetails); err != nil {
			log.WithError(err).Error("Failed to send request")
			continue
		}
	}
	return nil
}

func sendRequest(payload []byte, webhookDetails *db.WebhookDetails) error {
	req, err := http.NewRequest("POST", webhookDetails.Url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	for kHeader, vHeader := range webhookDetails.Headers {
		req.Header.Set(kHeader, vHeader)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(string(responseData))
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Status code != 200 while trying to send request, status: %v", resp.Status))
	}
	log.Infof("Message was sent succesfully: %s", webhookDetails.Url)
	return nil
}
