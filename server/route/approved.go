package route

import "net/http"

// Approved is a route for checking a list of all approved usernames.
// Intended for server-side mod to update the whitelist.
// !! REQUIRES API KEY.
func Approved(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
