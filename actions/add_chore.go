package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

// AddChore accepts a chorename and adds it to the Chores table
func AddChore(chorename string) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return err
	}

	defer conn.Close()

	_, err = conn.Exec(
		`INSERT INTO Chores(chorename) VALUES(?)`,
		chorename,
	)
	if err != nil {
		log.Printf("Error inserting chore: %s", err)
		return err
	}

	log.Printf("Chore added: %s", chorename)
	return nil
}
