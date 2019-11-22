package handlers

import (
	"bytes"
	"encoding/json"
	"komodorIntegrations/db"
	"net/http"
)

// SlackClient API for message sending
type SlackClient struct {
	Channel  string
	TokenURL string
}

// SendMessage sends slack message
func SendSlackMessage(message string, details db.Details) error {
	payload := map[string]string{
		"channel":  "integrations",
		"icon_url": slackbotIcon,
		"username": slackbotUsername,
		"text":     message,
	}

	payloadStr, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", details.WebhookUrl, bytes.NewBuffer([]byte(payloadStr)))
	if err != nil {
		return err
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
