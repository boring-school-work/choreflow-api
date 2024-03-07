package actions_test

import (
	"testing"

	"github.com/DaveSaah/choreflow-api/actions"
)

// func TestAddChore(t *testing.T) {
// 	if err := actions.AddChore("Go buy new slippers"); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestAssignChore(t *testing.T) {
// 	date_assign := time.Now().Format(time.DateOnly)
// 	date_due := time.Now().Format(time.DateOnly)
//
// 	assignment := &types.Assignment{
// 		Cid:         16,
// 		Sid:         4,
// 		DateAssign:  date_assign,
// 		DateDue:     date_due,
// 		WhoAssigned: 11,
// 	}
//
// 	if err := actions.AssignChore(10, assignment); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestDeleteChore(t *testing.T) {
// 	if err := actions.DeleteChore(18); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestEditChore(t *testing.T) {
// 	if err := actions.EditChore(16, "Go buy new slippers"); err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestLogin(t *testing.T) {
	person, err := actions.Login("someone@mail.com", "2xkWa24ABN")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Person: %v\n", person)
}
