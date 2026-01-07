package database

import (
	"errors"
	"log"
	"main/entity"
	"time"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	// Configure connection pool for better performance
	Connector.DB().SetMaxOpenConns(100)
	Connector.DB().SetMaxIdleConns(10)
	Connector.DB().SetConnMaxLifetime(time.Hour)

	// Test the connection
	if err := Connector.DB().Ping(); err != nil {
		return err
	}

	log.Println("Connection success")
	return nil
}

func Migrate(table *entity.Technique) error {
	if Connector == nil {
		return errors.New("database connection is nil")
	}

	result := Connector.AutoMigrate(&table)
	if result.Error != nil {
		return result.Error
	}

	log.Println("Table migrated")
	return nil
}
