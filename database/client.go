package database

import (
	"log"
	"main/entity"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	// Connector, err = gorm.Open("mysql", "root:judo-test-password@tcp(godockerDB)/techniques")
	Connector, err = gorm.Open("mysql", "sql11482611:ccFfrgmwy7@tcp(sql11.freemysqlhosting.net)/sql11482611")
	if err != nil {
		return err
	}

	log.Println("Connection success")
	return nil
}

func Migrate(table *entity.Technique) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
