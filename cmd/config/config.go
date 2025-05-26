package config

import (
	"log"

	"gorm.io/gorm"
)

var Config AppConfig
var DB *gorm.DB

func Load() {
	var err error

	Config, err = loadEnv()
	if err != nil {
		log.Fatal("Unable to load env file", err)
	}

	DB, err = initializeDB()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	InitializeValidator()
}
