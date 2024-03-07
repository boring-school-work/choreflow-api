package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

// AddChore adds a chore to the assignmentbase
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

	return err
}
