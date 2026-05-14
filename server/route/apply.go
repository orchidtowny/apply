package route

import (
	"apply/definition"
	"encoding/json"
	"net/http"
)

// Apply is a route for sending a new application to be whitelisted.
// Intended for new players to use in the frontend.
func Apply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body []byte
	_, readErr := r.Body.Read(body)
	if readErr != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	var jsonBody definition.ApplicationRequest
	marshalErr := json.Unmarshal(body, &jsonBody)
	if marshalErr != nil {
		http.Error(w, "Failed to parse body", http.StatusInternalServerError)
		return
	}

	// insert, validate
	// return created

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
