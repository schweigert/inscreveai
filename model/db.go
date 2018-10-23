package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Db() *gorm.DB {
	for {
		db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=rudy password=postgres sslmode=disable")
		if err != nil {
			log.Println("Cant connect to Postgresql:", err)
			log.Println("Sleep 5 seconds")
			time.Sleep(5 * time.Second)
		}

		db.LogMode(true)
		return db
	}
}
