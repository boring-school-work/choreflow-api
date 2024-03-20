package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/DaveSaah/choreflow-api/actions"
)

// addChore handles a GET request add a chore to the database.
// The request body is a string representing the chore name
func addChore(w http.ResponseWriter, r *http.Request) {
	errorResponse := response{msg: "Chore not added", code: http.StatusInternalServerError}
	successResponse := response{msg: "Chore added", code: http.StatusOK}

	chorename, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	chore := string(chorename)
	err = actions.AddChore(chore)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	successResponse.write(w)
}

// ChoreHandler handles a request to add, delete, or edit a chore.
func ChoreHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodPost:
		addChore(w, r)
	}
}
