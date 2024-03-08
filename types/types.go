// package types contains the entity definitions for the database
package types

type Assignment struct {
	DateDue       string
	LastUpdated   string
	DateCompleted string
	DateAssign    string
	Assignmentid  uint
	Cid           uint
	Sid           uint
	WhoAssigned   uint
}

type People struct {
	Passwd string
	Fname  string
	Lname  string
	Email  string
	Dob    string
	Tel    string
	Pid    uint
	Rid    uint
	Fid    uint
	Gender uint
}
