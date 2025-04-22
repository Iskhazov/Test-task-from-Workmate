package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json") // Tell the client we’re sending JSON
	w.WriteHeader(status)                              // Set HTTP status code
	return json.NewEncoder(w).Encode(v)                // Encode and send the response body as JSON
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
