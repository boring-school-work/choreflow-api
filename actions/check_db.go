package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

// CheckDB checks if the database is alive
func CheckDB() error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initializing db: %s", err)
		return err
	}

	defer conn.Close()

	return db.IsAlive(conn)
}
