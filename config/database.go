package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Driver for GORM
)

// DB Context
var DB *gorm.DB

// Database Initialization
func Database() *gorm.DB {
	driver := os.Getenv("DATABASE_DRIVER")
	database := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(driver, database)

	if err != nil {
		log.Panic(err)
	}
	log.Println("Database Connected")

	return DB
}
