package main

// Config holds the application configuration
var Config = struct {
	DbUri     string `env:"MONGO_URI" envDefault:"mongodb://mongodb-ha-europe-west1-b-1:27017,mongodb-ha-europe-west1-b-2:27017,mongodb-ha-europe-west1-b-3:27017/pxportal?replicaSet=rs0"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"json"`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"info"`
	Addr      string `env:"ADDR" envDefault:":8060"`

	SlackToken string `env:"SLACK_WEBHOOK" envDefault:"https://hooks.slack.com/services/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"`
}{}
