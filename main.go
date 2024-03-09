package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /choreflow/api/v1/add-chore", AddChoreHandler)
	mux.HandleFunc("GET /choreflow/api/v1/status", CheckDBHandler)
	mux.HandleFunc("POST /choreflow/api/v1/assign-chore/{pid}", AssignChoreHandler)

	log.Println("Starting api server at http://localhost:3211/choreflow/api/v1/")

	if err := http.ListenAndServe(":3211", mux); err != nil {
		log.Println(err)
	}
}
