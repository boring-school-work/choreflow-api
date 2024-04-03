package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DaveSaah/choreflow-api/actions"
)

// ChoresHandler handles a GET request to retrieve all chores from the database
func ChoresHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)
	errorResponse := response{msg: "Error retrieving chores", code: http.StatusInternalServerError}

	chores, err := actions.GetAllChores()
	if err != nil {
		log.Printf("Error retrieving chores: %s", err)

		if err == sql.ErrNoRows {
			errorResponse.update("No chores available")
		}

		errorResponse.write(w)
	}

	// build a json object with the chores slice
	w.Write([]byte(`{"chores": [`))
	for i, chore := range chores {

		log.Printf("Chore: %s", chore)
		log.Printf("i: %d", i)

		w.Write([]byte(`"` + chore + `"`))
		if i != len(chores)-1 {
			w.Write([]byte(","))
		}
	}

	w.Write([]byte(`]}`))
	w.WriteHeader(http.StatusOK)
}
