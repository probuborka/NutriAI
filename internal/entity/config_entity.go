package entity

const (
	// NutriAI
	Port = "8080"
	// GigaChat
	ApiKey = ""
	// Redis
	RedisHost = "localhost"
	RedisPort = "6379"
)

type HTTPConfig struct {
	Port string
}

type Api struct {
	Key string
}

type Redis struct {
	Host string
	Port string
}
