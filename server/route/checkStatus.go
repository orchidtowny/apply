package route

import "net/http"

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

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
