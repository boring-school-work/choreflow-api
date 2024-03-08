package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
	"github.com/DaveSaah/choreflow-api/types"
)

// assigns a chore to a child
func AssignChore(pid uint, assignment *types.Assignment) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return err
	}

	defer conn.Close()

	// begin transaction
	tx, err := conn.Begin()
	if err != nil {
		log.Printf("Cannot start db transaction: %s\n", err)
		return err
	}

	defer func() {
		_ = tx.Rollback() // aborts transaction
	}()

	// execute first transaction
	res, err := tx.Exec(
		`INSERT INTO Assignment(cid, sid, date_assign, date_due, who_assigned) 
    VALUES(?, ?, ?, ?, ?)`,
		assignment.Cid,
		assignment.Sid,
		assignment.DateAssign,
		assignment.DateDue,
		assignment.WhoAssigned,
	)
	if err != nil {
		log.Printf("Cannot execute db transaction: %s\n", err)
		return err
	}

	// get last inserted id
	assignmentid, _ := res.LastInsertId()

	// execute second transaction
	_, err = tx.Exec(
		`INSERT INTO Assigned_people(pid, assignmentid) 
    VALUES (?, ?)`,
		pid,
		assignmentid,
	)
	if err != nil {
		log.Printf("Cannot execute db transaction: %s\n", err)
		return err
	}

	// commit transaction
	if err = tx.Commit(); err != nil {
		log.Printf("Cannot commit db transaction: %s\n", err)
		return err
	}

	log.Printf("Chore assigned to: %d (pid)", pid)
	return nil
}
