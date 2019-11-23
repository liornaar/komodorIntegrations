package db

import (
	// Import this so we don't have to use qm.Limit etc.
	"database/sql"
	"encoding/json"
	"fmt"

	"komodorIntegrations/config"
	"sync"

	_ "github.com/lib/pq" // here
	_ "github.com/volatiletech/sqlboiler/queries/qm"
)

var doOnce sync.Once
var db *sql.DB
var err error

var Client *DbClient

type DbClient struct {
	connection *sql.DB
}

func NewDbClient() (*DbClient, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	Client = &DbClient{
		connection: conn,
	}
	return Client, Client.connection.Ping()
}

type Webhooks struct {
	Apps map[string]WebhookDetails `json:"apps"`
}

type WebhookDetails struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (db *DbClient) GetWebhooks(username string) (*Webhooks, error) {
	var details Details
	err := db.connection.QueryRow(fmt.Sprintf("SELECT webhooks from integrations_schema.integrations WHERE username = '%s'", username)).Scan(&details.Webhooks)
	var webhooksMap Webhooks
	err = json.Unmarshal(details.Webhooks, &webhooksMap)
	if err != nil {
		return nil, err
	}

	return &webhooksMap, nil
}

func (db *DbClient) Close() {
	db.connection.Close()
}
