package route

import "net/http"

// CreateInvite is a route for creating an invite.
// Intended only for the backend mod, players run the command, and it sends the request.
// !! REQUIRES API KEY
func CreateInvite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
