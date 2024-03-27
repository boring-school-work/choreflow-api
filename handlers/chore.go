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

// deleteChore handles a DELETE request to delete a chore from the database.
// The request path contains the chore id.
func deleteChore(w http.ResponseWriter, r *http.Request) {
	errorResponse := response{msg: "Chore not deleted", code: http.StatusInternalServerError}
	successResponse := response{msg: "Chore deleted", code: http.StatusOK}
	errorEmptyID := response{msg: "Chore ID is empty", code: http.StatusBadRequest}
	errorBadID := response{msg: "Chore ID is invalid", code: http.StatusBadRequest}

	id := r.PathValue("id")
	if id == "" {
		errorEmptyID.write(w)
		log.Println(errorEmptyID.msg)
		return
	}

	cid, err := strconv.Atoi(id)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	err = actions.DeleteChore(cid)
	if err != nil && err.Error() == "id not found" {
		errorBadID.write(w)
		return
	}

	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	successResponse.write(w)
}

func editChore(w http.ResponseWriter, r *http.Request) {
	errorResponse := response{msg: "Chore not edited", code: http.StatusInternalServerError}
	successResponse := response{msg: "Chore edited", code: http.StatusOK}
	errorEmptyID := response{msg: "Chore ID is empty", code: http.StatusBadRequest}
	errorBadID := response{msg: "Chore ID is invalid", code: http.StatusBadRequest}
	errorEmptyBody := response{msg: "Empty body", code: http.StatusBadRequest}

	id := r.PathValue("id")
	if id == "" {
		errorEmptyID.write(w)
		log.Println(errorEmptyID.msg)
		return
	}

	cid, err := strconv.Atoi(id)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	chorename, err := io.ReadAll(r.Body)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	chore := string(chorename)

	// check if an empty body was passed
	if chore == "" {
		errorEmptyBody.write(w)
		log.Println(errorEmptyBody.msg)
		return
	}

	err = actions.EditChore(cid, chore)
	if err != nil && err.Error() == "id not found" {
		errorBadID.write(w)
		return
	}

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
	case http.MethodDelete:
		deleteChore(w, r)
	case http.MethodPatch:
		editChore(w, r)
	}
}
