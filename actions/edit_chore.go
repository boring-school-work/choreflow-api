package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

func EditChore(cid uint, chorename string) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return err
	}

	defer conn.Close()

	_, err = conn.Exec(
		`UPDATE Chores SET chorename=? WHERE cid=?`,
		chorename,
		cid,
	)

	return err
}
