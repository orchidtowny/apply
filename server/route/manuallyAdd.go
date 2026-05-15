package route

import (
	"net/http"
)

// ManuallyAdd is a route for manually adding players to the list.
// Intended for server-side mod to add players who didn't apply but should be whitelisted.
// !! REQUIRES API KEY.
func ManuallyAdd(w http.ResponseWriter, r *http.Request) {
	/*
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

			w.Header().Set("Content-Type", "application/json")

			var usernames []string
			approved := util.GetApprovedApplications()
			for _, user := range approved {
				usernames = append(usernames, user.Username)
			}

			body, marshalErr := json.Marshal(definition.ApprovedUsers{
				Usernames: usernames,
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
	*/

	http.Error(w, "Not Implemented", http.StatusNotImplemented)
	return
}
