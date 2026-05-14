package route

import "net/http"

// Reject is a route for rejecting an application.
// Intended for server-side mod to reject applicants.
// !! REQUIRES API KEY.
func Reject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
