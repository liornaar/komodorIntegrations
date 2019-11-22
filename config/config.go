package config

// Config - Kafka
var Addr = String("ADDR", ":8000")
var LogFormat = String("LOG_FORMAT", "json")
var LogLevel = String("LOG_LEVEL", "info")
var DbHost = String("DB_HOST", "localhost")
var DbPort = String("DB_PORT", "5432")
var DbUser = String("DB_USER", "user")
var DbPass = String("DB_PASS", "passpass")
var DbName = String("DB_NAME", "komodor-main")
