package route

import (
	"applyServer/definition"
	"applyServer/util"
	"encoding/json"
	"net/http"
)

// Info is a route for getting information to show on the form.
// Intended only for frontend, getting rules, etc.
func Info(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	body, marshalErr := json.Marshal(definition.InfoResponse{
		Open:  util.Config.Open,
		Rules: util.Config.Rules,
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
