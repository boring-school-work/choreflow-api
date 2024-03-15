package main

import (
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// define base url and port
var (
	baseURL string = "/choreflow/api/v1/"
	port    string = ":3211"
)

// getURI returns a string that represents the HTTP method and path
func getURI(method, endpoint string) string {
	return method + " " + baseURL + endpoint
}

// @title Choreflow API
// @version 1.0
// @description A API for choreflow: https://github/DaveSaah/choreflow
// @termsOfService http://swagger.io/terms/

// @contact.name David Saah
// @contact.url https://github.com/DaveSaah
// @contact.email saahdavid17@gmail.com

// @license.name MIT
// @license.url https://github.com/DaveSaah/choreflow-api/LICENSE

// @host petstore.swagger.io
// @BasePath /choreflow/api/v1/
func main() {
	mux := http.NewServeMux()

	// serve swagger docs
	mux.HandleFunc(getURI("GET", "swagger/swagger.json"), SwaggerSpecHandler)

	// serve Swagger UI
	mux.HandleFunc(
		getURI("GET", "swagger/ui/"),
		httpSwagger.Handler(httpSwagger.URL("http://localhost:3211/choreflow/api/v1/swagger/swagger.json")),
	)

	// serve API endpoints
	mux.HandleFunc(getURI("POST", "add-chore"), AddChoreHandler)
	mux.HandleFunc(getURI("GET", "status"), CheckDBHandler)
	mux.HandleFunc(getURI("POST", "assign-chore/{pid}"), AssignChoreHandler)

	log.Printf("Starting api server at http://localhost%s%s\n", port, baseURL)
	log.Printf("Starting api docs at http://localhost%s%s%s\n", port, baseURL, "swagger/ui")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Println(err)
	}
}
