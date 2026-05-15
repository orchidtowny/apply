package route

import (
	"applyServer/definition"
	"applyServer/util"
	"encoding/json"
	"net/http"
	"strings"
)

// Approved is a route for checking a list of all approved usernames.
// Intended for server-side mod to update the whitelist.
// !! REQUIRES API KEY.
func Approved(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}
	if !strings.HasPrefix(authorization, "Bearer ") {
		http.Error(w, "Bearer authorization required", http.StatusUnauthorized)
		return
	}

	if util.Config.ApiKey != strings.TrimPrefix(authorization, "Bearer ") {
		http.Error(w, "Invalid authorization", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var usernames []definition.MinecraftPlayer
	approved := util.GetApprovedApplications()
	for _, user := range approved {
		usernames = append(usernames, definition.MinecraftPlayer{
			Uuid:     user.Uuid,
			Username: user.Username,
		})
	}

	body, marshalErr := json.Marshal(usernames)
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
