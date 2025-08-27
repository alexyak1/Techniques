package database

import (
	"log"
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

	// Tune connection pool to reduce unnecessary work
	sqlDB := Connector.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Connection success")
	return nil
}

func Migrate(models ...interface{}) {
	Connector.AutoMigrate(models...)
	log.Println("Tables migrated")
}
