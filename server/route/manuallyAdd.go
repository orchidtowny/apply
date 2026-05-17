package route

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"orchidmc.org/apply/server/definition"
	"orchidmc.org/apply/server/util"

	"github.com/google/uuid"
)

// ManuallyAdd is a route for manually adding players to the list.
// Intended for server-side mod to add players who didn't apply but should be whitelisted.
// !! REQUIRES API KEY.
func ManuallyAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	var jsonBody definition.ManuallyAddRequest
	marshalErr := json.Unmarshal(body, &jsonBody)
	if marshalErr != nil {
		http.Error(w, "Failed to parse body: "+marshalErr.Error(), http.StatusInternalServerError)
		return
	}

	if jsonBody.Username == "" {
		http.Error(w, "\"username\" must not be blank", http.StatusBadRequest)
		return
	}

	newApplicationId := uuid.New().String()
	foundUuid := util.GetPlayerUuid(jsonBody.Username)

	if foundUuid == "" {
		http.Error(w, "Couldn't fetch player information from Mojang.", http.StatusFailedDependency)
		return
	}

	util.CreateApplication(definition.Application{
		Id: newApplicationId,

		Username:                 jsonBody.Username,
		Uuid:                     foundUuid,
		Age:                      0,
		WhereDidYouFindTheServer: "",
		Bio:                      "",

		Status: 2,
	})

	w.Header().Set("Content-Type", "application/json")

	responseBody, marshalErr := json.Marshal(definition.GenericPostResponse{
		Success: true,
		Message: "Added and approved " + jsonBody.Username,
	})
	if marshalErr != nil {
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(responseBody)
	if writeErr != nil {
		http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		return
	}

	return
}
