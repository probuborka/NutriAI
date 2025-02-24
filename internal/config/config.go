package config

import "github.com/probuborka/NutriAI/internal/entity"

type Config struct {
	HTTP entity.HTTPConfig
	// DB   entityconfig.DBConfig
	// Auth entityconfig.Authentication
}

func New() (*Config, error) {
	// // password
	// password := os.Getenv("TODO_PASSWORD")

	// //port
	// port := os.Getenv("TODO_PORT")
	// if port == "" {
	// 	port = entityconfig.Port
	// }
	port := entity.Port

	// //db
	// dbFile := os.Getenv("TODO_DBFILE")
	// if dbFile == "" {
	// 	dbFile = filepath.Join(entityconfig.DBDir, "/", entityconfig.DBName)
	// }
	// dbDriver := entityconfig.DBDriver
	// dbCreate := entityconfig.DBCreate

	return &Config{
		HTTP: entity.HTTPConfig{Port: port},
		// DB: entityconfig.DBConfig{
		// 	Driver: dbDriver,
		// 	File:   dbFile,
		// 	Create: dbCreate,
		// },
		// Auth: entityconfig.Authentication{
		// 	Password: password,
		// },
	}, nil
}
