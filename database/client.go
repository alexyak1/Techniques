package database

import (
	"log"
	"main/entity"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Migrate(table *entity.Technique) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
