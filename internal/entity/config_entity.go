package entity

const (
	// NutriAI
	Port = "8090"
	// GigaChat
	ApiKey = "ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOmEwN2Q5YjhkLWVlNDAtNDUzZS04MTk1LTYzMDQxODU0NjYwMA=="
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
