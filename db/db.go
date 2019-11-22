package db

import (
	// Import this so we don't have to use qm.Limit etc.
	"database/sql"
	"fmt"

	"komodorIntegrations/config"
	"sync"

	_ "github.com/lib/pq" // here
	_ "github.com/volatiletech/sqlboiler/queries/qm"
)

var doOnce sync.Once
var db *sql.DB
var err error

var dbClient *DbClient

type DbClient struct {
	connection *sql.DB
}

func NewDbClient() (*DbClient, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	dbClient = &DbClient{
		connection: conn,
	}
	return dbClient, dbClient.connection.Ping()
}

func (db *DbClient) GetUserDetails(username string) (*Details, error) {
	return nil, nil
}

func (db *DbClient) Close() {
	db.connection.Close()
}
