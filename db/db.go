package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Init initializes a new connection to the database
func Init() (*sql.DB, error) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("DBHOST"), os.Getenv("DBPORT"))
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  time.UTC,
	}

	return sql.Open("mysql", cfg.FormatDSN())
}

func IsAlive(db *sql.DB) error {
	return db.Ping()
}
