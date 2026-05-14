package route

import "net/http"

// Approve is a route for approving an application.
// Intended for server-side mod to approve applicants.
// !! REQUIRES API KEY.
func Approve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
