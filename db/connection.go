package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const DSN = "host=localhost user=golang password=golang dbname=db_restapi_go port=5432 sslmode=disable TimeZone=America/Bogota"

var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}

}
