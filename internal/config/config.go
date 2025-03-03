package config

import (
	"os"

	"github.com/probuborka/NutriAI/internal/entity"
)

type Config struct {
	HTTP  entity.HTTPConfig
	Api   entity.Api
	Redis entity.Redis
	Log   entity.Log
}

func New() (*Config, error) {
	//port
	port := os.Getenv("NUTRIAI_PORT")
	if port == "" {
		port = entity.Port
	}

	//API_KEY
	key := os.Getenv("API_KEY")
	if key == "" {
		key = entity.ApiKey
	}

	//RedisHost
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = entity.RedisHost
	}

	//RedisPort
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = entity.RedisPort
	}

	//
	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = entity.LogFile
	}

	return &Config{
		HTTP: entity.HTTPConfig{
			Port: port,
		},
		Api: entity.Api{
			Key: key,
		},
		Redis: entity.Redis{
			Host: redisHost,
			Port: redisPort,
		},
		Log: entity.Log{
			File: logFile,
		},
	}, nil
}
