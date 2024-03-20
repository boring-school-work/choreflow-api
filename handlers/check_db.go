package handlers

import (
	"log"
	"net/http"

	"github.com/DaveSaah/choreflow-api/actions"
)

// CheckDBHandler handles a GET request to check if the
// database is alive. It returns a json response with the
// status of the database (up or down)
func CheckDBHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)

	// define responses
	errorResponse := response{msg: "DB is down", code: http.StatusInternalServerError}
	successResponse := response{msg: "DB is up", code: http.StatusOK}

	err := actions.CheckDB()
	if err != nil {
		errorResponse.write(w)
		log.Println(err)
		return
	}

	successResponse.write(w)
}
