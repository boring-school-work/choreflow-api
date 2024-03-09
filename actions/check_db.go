package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

func CheckDB() error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Error initializing db: %s", err)
		return err
	}

	defer conn.Close()

	return db.IsAlive(conn)
}
