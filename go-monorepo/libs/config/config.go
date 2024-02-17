package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type sharedConfig struct {
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASSWORD string

	AWS_ACCESS_KEY string
	AWS_SECRET_KEY string
	AWS_REGION     string
}

type serverConfig struct {
	HOST string
	PORT string
}

var SharedConfig sharedConfig
var ApigatewayConfig serverConfig
var IndexerConfig serverConfig
var TradingConfig serverConfig
var SiteConfig serverConfig

func getEnv(key string, defaultValue string) string {
	val := os.Getenv("REDIS_HOST")
	if val == "" {
		return defaultValue
	}
	return val
}

func InitializeConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SharedConfig = sharedConfig{
		REDIS_HOST:     getEnv("REDIS_HOST", "localhost"),
		REDIS_PASSWORD: getEnv("REDIS_PASSWORD", ""),
		REDIS_PORT:     getEnv("REDIS_PORT", "6379"),

		AWS_ACCESS_KEY: getEnv("AWS_ACCESS_KEY", ""),
		AWS_SECRET_KEY: getEnv("AWS_SECRET_KEY", ""),
		AWS_REGION:     getEnv("AWS_REGION", ""),
	}

	ApigatewayConfig = serverConfig{
		HOST: getEnv("API_GATEWAY_HOST", "localhost"),
		PORT: getEnv("API_GATEWAY_PORT", "3001"),
	}

	SiteConfig = serverConfig{
		HOST: getEnv("SITE_HOST", "localhost"),
		PORT: getEnv("SITE_PORT", "4001"),
	}

	IndexerConfig = serverConfig{
		HOST: getEnv("INDEXER_HOST", "localhost"),
		PORT: getEnv("INDEXER_PORT", "4010"),
	}

	TradingConfig = serverConfig{
		HOST: getEnv("TRADING_HOST", "localhost"),
		PORT: getEnv("TRADING_PORT", "4020"),
	}

}
