package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

func DeleteChore(cid uint) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return err
	}

	defer conn.Close()

	_, err = conn.Exec(`DELETE FROM Chores WHERE cid=?`, cid)
	if err != nil {
		log.Printf("Cannot delete chore: %s\n", err)
		return err
	}

	log.Printf("Chore deleted: %d (cid)", cid)
	return nil
}
