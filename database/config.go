package database

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var GetConnectionString = func(config Config) string {
	// connectionString := "root:judo-test-password@tcp(godockerDB)/techniques"
	connectionString := "sql11482611:ccFfrgmwy7@tcp(sql11.freemysqlhosting.net)/sql11482611"

	return connectionString
}
