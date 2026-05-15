package route

import (
	"apply/definition"
	"apply/util"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Approve is a route for approving an application.
// Intended for server-side mod to approve applicants.
// !! REQUIRES API KEY.
func Approve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
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

	application := util.GetApplication(id)
	if application == nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	util.ApproveApplication(application.Id)

	body, marshalErr := json.Marshal(definition.GenericPostResponse{
		Success: true,
		Message: "Approved application for " + application.Username,
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

	go func() {
		util.SendWebhookMessage(definition.DiscordWebhook{
			Username: "Applications",
			Embeds: []definition.DiscordWebhookEmbed{
				{
					Title:     "Approved Application for " + application.Username,
					Timestamp: time.Now().Format(time.RFC3339),
					Fields: []definition.DiscordWebhookEmbedField{
						{
							Name:   "ID",
							Value:  id,
							Inline: true,
						},
					},
					Author: &definition.DiscordWebhookEmbedAuthor{
						Name:    application.Username,
						IconUrl: "https://mc-heads.net/head/" + application.Username,
					},
				},
			},
		})
	}()

	return
}
