package database

import "fmt"

type Config struct {
	ServerName string
	User string
	Password string
	DB string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}