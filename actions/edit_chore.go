package actions

import (
	"errors"
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

// EditChore edits a chore in the database
// Arguments:
//
//	cid -> the chore id
func EditChore(cid int, chorename string) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return err
	}

	defer conn.Close()

	res, err := conn.Exec(
		`UPDATE Chores SET chorename=? WHERE cid=?`,
		chorename,
		cid,
	)
	if err != nil {
		log.Printf("Cannot edit chore: %s\n", err)
		return err
	}

	count, _ := res.RowsAffected()

	if count == 0 {
		log.Printf("No chore found with id: %d\n", cid)
		return errors.New("id not found")
	}

	log.Printf("Chore edited: %d (cid)", cid)
	return nil
}
