package handlers

import (
	"net/http"
)

// response struct to handle the http response body and status code
type response struct {
	msg  string
	code int
}

// write writes the response to the http.ResponseWriter.
// It sets the content type to application/json and writes
// the response body as a json object and sets the status code
func (res *response) write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.code)
	w.Write([]byte(`{"msg": "` + res.msg + `"}`))
}
