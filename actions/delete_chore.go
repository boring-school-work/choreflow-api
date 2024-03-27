package actions

import (
	"errors"
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

// DeleteChore deletes a chore from the database
// Arguments:
//
//	cid -> the chore id
func DeleteChore(cid int) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM Chores WHERE cid=?`, cid)
	if err != nil {
		log.Printf("Cannot delete chore: %s\n", err)
		return err
	}

	count, _ := res.RowsAffected()

	if count == 0 {
		log.Printf("No chore found with id: %d\n", cid)
		return errors.New("id not found")
	}

	log.Printf("Chore deleted: %d (cid)", cid)
	return nil
}
