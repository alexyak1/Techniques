package database

type Config struct {
	ServerName string
	User string
	Password string
	DB string
}

var GetConnectionString = func(config Config) string {
	connectionString := "root:judo-test-password@tcp(godockerDB)/techniques"

	return connectionString
}