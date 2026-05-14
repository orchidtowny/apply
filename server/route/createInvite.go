package route

import (
	"apply/util"
	"net/http"
	"strings"
)

// CreateInvite is a route for creating an invite.
// Intended only for the backend mod, players run the command, and it sends the request.
// !! REQUIRES API KEY
func CreateInvite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}
	if strings.HasPrefix(authorization, "Bearer ") {
		http.Error(w, "Bearer authorization required", http.StatusUnauthorized)
		return
	}

	if util.Config.ApiKey != strings.TrimPrefix(authorization, "Bearer ") {
		http.Error(w, "Invalid authorization", http.StatusForbidden)
		return
	}

	// read body, { username: String }
	// create invite code: uuid, creator: username

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
