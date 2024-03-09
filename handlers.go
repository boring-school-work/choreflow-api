package main

import (
	"log"
	"net/http"

	"github.com/DaveSaah/choreflow-api/actions"
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

