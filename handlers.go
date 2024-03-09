package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/DaveSaah/choreflow-api/actions"
	"github.com/DaveSaah/choreflow-api/types"
)
func CheckDBHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)

	err := actions.CheckDB()
	if err != nil {
		http.Error(w, "DB is down", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write([]byte("DB is running"))
}

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
