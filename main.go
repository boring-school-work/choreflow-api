package main

import (
	"log"
	"net/http"

	"github.com/DaveSaah/choreflow-api/handlers"
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

func main() {
	mux := http.NewServeMux()

	// serve swagger docs
	mux.HandleFunc(getURI("GET", "swagger/swagger.yaml"), handlers.SwaggerSpecHandler)

	// serve Swagger UI
	mux.HandleFunc(
		getURI("GET", "swagger/ui/"),
		httpSwagger.Handler(httpSwagger.URL("http://localhost:3211/choreflow/api/v1/swagger/swagger.yaml")),
	)

	// serve API endpoints
	mux.HandleFunc(getURI("GET", "status"), handlers.CheckDBHandler)
	mux.HandleFunc(getURI("", "chore/{id}"), handlers.ChoreHandler)
	mux.HandleFunc(getURI("", "chore"), handlers.ChoreHandler)
	mux.HandleFunc(getURI("", "assignment/{pid}"), handlers.AssignmentHandler)

	log.Printf("Starting api server at http://localhost%s%s\n", port, baseURL)
	log.Printf("Starting api docs at http://localhost%s%s%s\n", port, baseURL, "swagger/ui")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Println(err)
	}
}
