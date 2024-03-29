package database

import (
	"log"
	"main/entity"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
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
