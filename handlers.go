package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/DaveSaah/choreflow-api/actions"
	"github.com/DaveSaah/choreflow-api/types"
)

// AddChoreHandler handles a GET request to add a chore.
// The request body is a string that represents the chore name
func AddChoreHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)

	chorename, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Chore not added", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	chore := string(chorename)
	err = actions.AddChore(chore)
	if err != nil {
		http.Error(w, "Chore not added", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write([]byte("Chore added"))
}

// CheckDBHandler handles a GET request to check if the
// database is alive. It returns a json response with the
// status of the database (up or down)
func CheckDBHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	err := actions.CheckDB()
	if err != nil {
		http.Error(w, "DB is down", http.StatusInternalServerError)
		log.Println(err)
		w.Write([]byte(`{"status": "down"}`))
		return
	}

	// write json response
	w.Write([]byte(`{"status": "up"}`))
}

// AssignChoreHandler handles a POST request to assign a chore to a user
// The request body has a pid url parameter and accepts a json that
// contains the assignment information
func AssignChoreHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)

	var assignment types.Assignment // initialise assignment type
	pid_str := r.PathValue("pid")

	// convert to pid from string to int
	pid, err := strconv.ParseUint(pid_str, 10, 64)
	if err != nil {
		http.Error(w, "Chore assignment failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// decode the request body into the assignment type
	err = json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		http.Error(w, "Chore assignment failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = actions.AssignChore(pid, &assignment)
	if err != nil {
		http.Error(w, "Chore assignment failed", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write([]byte("Chore assigned"))
}

// SwaggerSpecHandler handles a GET request to serve the swagger config file
// for the api's documentation
func SwaggerSpecHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)
	http.ServeFile(w, r, "./docs/swagger.json")
}
