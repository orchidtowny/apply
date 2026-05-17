package route

import (
	"encoding/json"
	"net/http"

	"orchidmc.org/apply/server/definition"
	"orchidmc.org/apply/server/util"
)

// CheckStatus is a route for checking if a username is approved.
// Intended for backend if the list doesn't contain a username yet OR if a user wants to check their status.
func CheckStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	application := util.GetApplicationByUsername(username)
	if application == nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	body, marshalErr := json.Marshal(definition.StatusResponse{
		Username: application.Username,
		Status:   application.Status,
	})
	if marshalErr != nil {
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(body)
	if writeErr != nil {
		http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		return
	}

	return
}
