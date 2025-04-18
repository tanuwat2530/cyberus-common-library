package utilities

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON sends a JSON response
func ResponseWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// RespondWithError sends an error response
func ResponseWithError(w http.ResponseWriter, status int, message string) {
	ResponseWithJSON(w, status, map[string]string{"error": message})
}

func ResponseWithTest(w http.ResponseWriter, message string) {
	//response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(message))
}
