package actions

import (
	"database/sql"
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

// GetChore retrieves a chore from the database.
//
// Arguments:
//
//	cid -> the chore id
//
//	Returns:
//	error -> an error if the chore is not found or the id is invalid
func GetChore(cid int) (string, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return "", err
	}

	defer conn.Close()

	var chorename string

	err = conn.QueryRow(
		`SELECT chorename FROM Chores WHERE cid = ?`,
		cid,
	).Scan(&chorename)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("Cannot retrieve chore: %s\n", err)
		return "", err
	case err != nil:
		log.Printf("Error querying the database: %s", err)
		return "", err
	}

	log.Printf("Chore retrieved: %d (cid)", cid)
	return chorename, nil
}
