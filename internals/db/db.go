package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDb() {
	cfg := mysql.Config{
		dbHost: os.Getenv("DB_HOST"),
		dbPort: os.Getenv("DB_ENV"),
		dbUser: os.Getenv("DB_USER"),
		dbPass: os.Getenv("DB_PASSWORD"),
		dbName: os.Getenv("DB_NAME"),
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
}
