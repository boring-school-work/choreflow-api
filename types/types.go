// package types contains the entity definitions for the database
package types

// Assignment struct contains the fields for the assignment entity
type Assignment struct {
	DateDue       string `json:"date-due"`
	LastUpdated   string `json:"last-updated"`
	DateCompleted string `json:"date-completed"`
	DateAssign    string `json:"date-assign"`
	Assignmentid  uint   `json:"assignmentid"`
	Cid           uint   `json:"cid"`
	Sid           uint   `json:"sid"`
	WhoAssigned   uint   `json:"who-assigned"`
}

// Person struct contains the fields for the person entity
type People struct {
	Passwd string `json:"passwd"`
	Fname  string `json:"fname"`
	Lname  string `json:"lname"`
	Email  string `json:"email"`
	Dob    string `json:"dob"`
	Tel    string `json:"tel"`
	Pid    uint   `json:"pid"`
	Rid    uint   `json:"rid"`
	Fid    uint   `json:"fid"`
	Gender uint   `json:"gender"`
}
