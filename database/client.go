package database

import (
	"log"
	"main/entity"

	// "github.com/jinzhu/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=ec2-34-248-169-69.eu-west-1.compute.amazonaws.com user=llxgckjvfqvhuw password=0d63932135f5c6a82950d2423b763fbc7308a12555ac31154ea968ebb86ac3ce dbname=d6juhn5dera7pp port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                                                                                                                          // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Connection success")
	log.Println(db)
	return nil
}

func Migrate(table *entity.Technique) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
