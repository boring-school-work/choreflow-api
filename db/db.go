package db

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Init initializes a new connection to the database
func Init() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "chores_mgt",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.UTC,
	}

	return sql.Open("mysql", cfg.FormatDSN())
}

func IsAlive(db *sql.DB) error {
	return db.Ping()
}
