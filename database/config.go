package database

type Config struct {
	ServerName string
	User       string
	Hash       string
	DB         string
}

var GetConnectionString = func(config Config) string {
	return config.User + ":" + config.Hash + "@tcp(" + config.ServerName + ")/" + config.User

}
