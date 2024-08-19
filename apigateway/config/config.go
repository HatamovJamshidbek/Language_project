package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_SERVICE     string
	LEARNING_SERVICE string
	SIGNING_KEY      string
	HTTP_PORT        string
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{}
	config.HTTP_PORT = cast.ToString(Coalesce("HTTP_PORT", "8080"))
	config.AUTH_SERVICE = cast.ToString(Coalesce("AUTH_SERVICE", "8075"))
	config.LEARNING_SERVICE = cast.ToString(Coalesce("LEARNING_SERVICE", "8070"))
	config.SIGNING_KEY = cast.ToString(Coalesce("SIGNING_KEY", ":secret"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
