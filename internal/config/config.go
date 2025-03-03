package config

import (
	"os"

	"github.com/probuborka/NutriAI/internal/entity"
)

type Config struct {
	HTTP  entity.HTTPConfig
	Api   entity.Api
	Redis entity.Redis
	// DB   entityconfig.DBConfig
	// Auth entityconfig.Authentication
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
	}, nil
}
