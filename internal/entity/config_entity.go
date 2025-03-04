package entity

const (
	//nutriAI
	Port = "8080"

	//gigaChat
	ApiKey = "ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOmEwN2Q5YjhkLWVlNDAtNDUzZS04MTk1LTYzMDQxODU0NjYwMA=="

	//redis
	RedisHost = "localhost"
	RedisPort = "6379"

	//log file
	LogFile = "./var/log/app.log"
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

type Log struct {
	File string
}
