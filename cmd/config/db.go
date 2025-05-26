package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeDB() (*gorm.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Config.DB_HOST,
		Config.DB_PORT,
		Config.DB_USER,
		Config.DB_PASSWORD,
		Config.DB_NAME,
		Config.DB_SSL_MODE,
	)
	return gorm.Open(postgres.Open(connStr), &gorm.Config{})
}
