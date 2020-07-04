package utils

import (
	"encoding/json"
	"net/http"
)

// Response is a struct to define body of http response
type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewResponse is a function to initialize a new Response{} struct for http response
func NewResponse(hasError bool, message string, data interface{}) Response {
	return Response{Error: hasError, Message: message, Data: data}
}

// Respond is a function to generate a response http with success, including a return in JSON format or not
func Respond(w http.ResponseWriter, httStatus int, response Response) {
	w.WriteHeader(httStatus)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}
