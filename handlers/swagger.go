package handlers

import (
	"log"
	"net/http"
)

// SwaggerSpecHandler handles a GET request to serve the swagger config file
func SwaggerSpecHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s to %s %s", r.Host, r.Method, r.URL.Path)
	http.ServeFile(w, r, "./docs/swagger.yaml")
}
