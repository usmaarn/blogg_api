package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initializeDB() (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Config.DB_HOST,
		Config.DB_PORT,
		Config.DB_USER,
		Config.DB_PASSWORD,
		Config.DB_NAME,
		Config.DB_SSL_MODE,
	)
	return sqlx.Connect("postgres", connStr)
}
