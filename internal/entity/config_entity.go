package entity

const (
	//nutriAI
	Port = "8080"

	//gigaChat
	ApiKey = "ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOjVlMmM3OWYxLTUwNDQtNDRkNi05NTY1LTA3NzBlNTkyMWNmMQ=="

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
