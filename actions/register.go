package actions

import (
	"log"

	"github.com/DaveSaah/choreflow-api/db"
	"github.com/DaveSaah/choreflow-api/types"
	"golang.org/x/crypto/bcrypt"
)

// EmailExistsError is an error type for when a user tries to register
// with an email that already exists
type EmailExistsError struct{}

// Error returns the error message for EmailExistsError type
func (e *EmailExistsError) Error() string {
	return "Email already exists"
}

// getHashPassword returns a hashed password using bcrypt and default cost
func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Unable to hash password: %s", err)
		return "", err
	}

	return string(hash), nil
}

// Register registers a user into the database
func Register(person *types.People) error {
	conn, err := db.Init()
	if err != nil {
		log.Printf("Cannot create the db connection: %s\n", err)
		return err
	}

	defer conn.Close()

	var count int

	// check if email already exists
	err = conn.QueryRow(
		`SELECT COUNT(email) FROM People WHERE email=?`,
		person.Email,
	).Scan(
		&count,
	)

	switch {
	case err != nil:
		log.Printf("Error querying the database: %s", err)
		return err

	case count > 0:
		log.Printf("Email already exists: %s", person.Email)
		return &EmailExistsError{}
	}

	// hash the password
	person.Passwd, err = getHashPassword(person.Passwd)
	if err != nil {
		log.Println(err)
		return err
	}

	// pick rid
	if person.Fid == 1 || person.Fid == 2 {
		person.Rid = person.Fid
	} else {
		person.Rid = 3
	}

	_, err = conn.Exec(
		`INSERT INTO 
    People(fname, lname, gender, dob, tel, email, passwd, rid, fid) 
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		person.Fname,
		person.Lname,
		person.Gender,
		person.Dob,
		person.Tel,
		person.Email,
		person.Passwd,
		person.Rid,
		person.Fid,
	)
	if err != nil {
		log.Printf("Error inserting into the database: %s", err)
		return err
	}

	log.Printf("User registered: %s", person.Email)
	return nil
}
