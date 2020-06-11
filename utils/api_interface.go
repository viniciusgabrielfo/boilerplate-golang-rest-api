package utils

import (
	"encoding/json"
	"net/http"
)

// Response is a function to generate a response http to API, in JSON format
func Response(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
