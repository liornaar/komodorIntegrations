package db

import "encoding/json"

type Details struct {
	ID                string          `json:"id"`
	Username          string          `json:"username"`
	SlackWebhookUrl   string          `json:"slack_url"`
	GenericWebhookUrl string          `json:"generic_webhook_url"`
	CreatedAt         string          `json:"created_at"`
	Webhooks          json.RawMessage `json:"webhooks"`
}
