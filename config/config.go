package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort    string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No.. .env file ")
	}
	config := Config{}
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7053"))
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}