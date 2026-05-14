package route

import "net/http"

// Apply is a route for sending a new application to be whitelisted.
// Intended for new players to use in the frontend.
func Apply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
