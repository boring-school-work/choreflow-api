package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DaveSaah/choreflow-api/actions"
	"github.com/DaveSaah/choreflow-api/types"
)

// AssignmentHandler handles a POST request to assign a chore to a user
// The request body has a pid url parameter and accepts a json that
// contains the assignment information
func AssignmentHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)

	// define responses
	errorResponse := response{msg: "Chore assignment failed", code: http.StatusInternalServerError}
	successResponse := response{msg: "Chore assigned", code: http.StatusOK}

	var assignment types.Assignment // initialise assignment type
	pid_str := r.PathValue("pid")

	// convert to pid from string to int
	pid, err := strconv.ParseUint(pid_str, 10, 64)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	// decode the request body into the assignment type
	err = json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	err = actions.AssignChore(pid, &assignment)
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	successResponse.write(w)
}
