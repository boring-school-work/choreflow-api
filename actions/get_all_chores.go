package actions

import (
	"database/sql"
	"log"

	"github.com/DaveSaah/choreflow-api/db"
)

// GetAllChores retrieves a chore from the database.
//
//	Returns:
//	[]string -> a slice of strings representing all chores
//	error -> an error if the chore is not found or the id is invalid
func GetAllChores() ([]string, error) {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return nil, err
	}

	defer conn.Close()

	var chores []string

	rows, err := conn.Query("SELECT chorename FROM Chores")

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No chores have been added: %s\n", err)
		return []string{"No chores available"}, err
	case err != nil:
		log.Printf("Error querying the database: %s", err)
		return nil, err
	}

	for rows.Next() {
		var chore string

		err = rows.Scan(&chore)
		if err != nil {
			log.Printf("Error scanning rows: %s", err)
		}

		chores = append(chores, chore)
	}

	log.Printf("All chores retrieved")
	log.Printf("Chores: %v", chores)
	return chores, nil
}
