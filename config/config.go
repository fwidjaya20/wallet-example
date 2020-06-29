package config

import "os"

// noinspection ALL
const (
	ENV 			= "ENV"
	ENV_DEVELOPMENT = "development"
	ENV_STAGING 	= "staging"
	ENV_PRODUCTION 	= "production"
)

// noinspection ALL
const (
	SERVICE_NAME string = "SERVICE_NAME"

	HTTP_ADDR = "HTTP_ADDR"

	DB_DRIVER = "DB_DRIVER"
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
	DB_NAME = "DB_NAME"

	MIGRATION_PATH = "MIGRATION_PATH"
)

var defaultConfig = map[string]string{
	// Common Configuration
	ENV:          ENV_DEVELOPMENT,
	SERVICE_NAME: "wallet-example",

	// Database Configuration
	DB_DRIVER: "postgres",
	DB_HOST:   "localhost",
	DB_PORT:   "5432",
	DB_NAME:   "wallet_example",
	DB_USER:   "postgres",
	DB_PASS:   "postgres",

	// Migration and Seeder
	MIGRATION_PATH: "internal/databases/migrations",

	// Transport
	HTTP_ADDR: ":8001",
}

func GetEnv(key string) string {
	r := os.Getenv(key)

	if r == "" {
		if _, ok := defaultConfig[key]; !ok {
			return ""
		}
		r = defaultConfig[key]
	}

	return r
}