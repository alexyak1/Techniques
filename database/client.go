package database

import (
	"database/sql"
	"log"
	"main/entity"
	"os"

	// "github.com/jinzhu/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Connector *gorm.DB

func Connect() error {
	var err error
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection success")
	log.Println(db)
	return nil
}

func Migrate(table *entity.Technique) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
