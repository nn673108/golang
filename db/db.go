package db

import (
	"fmt"
	"gorm.io/driver/postgres" // import gorm postgres
	"gorm.io/gorm"            // import gorm
	"log"
	"os"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := os.Getenv("POSTGRES_DNS")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	fmt.Println("Connect Database Successfuly!")
}
