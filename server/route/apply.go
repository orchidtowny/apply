package route

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"strconv"
	"time"

	"orchidmc.org/apply/server/definition"
	"orchidmc.org/apply/server/util"

	"github.com/google/uuid"
)

// Apply is a route for sending a new application to be whitelisted.
// Intended for new players to use in the frontend.
func Apply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	var jsonBody definition.ApplicationRequest
	marshalErr := json.Unmarshal(body, &jsonBody)
	if marshalErr != nil {
		http.Error(w, "Failed to parse body: "+marshalErr.Error(), http.StatusInternalServerError)
		return
	}

	// Validation

	if jsonBody.Username == "" {
		http.Error(w, "\"username\" must not be blank", http.StatusBadRequest)
		return
	}

	if jsonBody.Age <= 0 || jsonBody.Age >= 100 {
		http.Error(w, "\"age\" invalid", http.StatusBadRequest)
		return
	}

	if jsonBody.WhereDidYouFindTheServer == "" {
		http.Error(w, "\"where_did_you_find_the_server\" must not be blank", http.StatusBadRequest)
		return
	}

	if jsonBody.Bio == "" {
		http.Error(w, "\"bio\" must not be blank", http.StatusBadRequest)
		return
	}

	if util.GetApplicationByUsername(jsonBody.Username) != nil {
		http.Error(w, "An application already exists!", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	fmt.Println("Creating application " + id + " for user " + jsonBody.Username)

	escapedUsername := html.EscapeString(jsonBody.Username)
	foundUuid := util.GetPlayerUuid(escapedUsername)

	if foundUuid == "" {
		http.Error(w, "Couldn't fetch your information from Mojang.", http.StatusFailedDependency)
		return
	}

	util.CreateApplication(definition.Application{
		Id: id,

		Username:                 escapedUsername,
		Uuid:                     foundUuid,
		Age:                      jsonBody.Age,
		WhereDidYouFindTheServer: html.EscapeString(jsonBody.WhereDidYouFindTheServer),
		Bio:                      html.EscapeString(jsonBody.Bio),

		Status: 0,
	})

	// Notification

	go func() {
		if util.Config.DiscordWebhookUrl == "" {
			fmt.Println("Not attempting webhook send")
			return
		}

		fmt.Println("Attempting webhook send")
		application := util.GetApplication(id)

		util.SendWebhookMessage(definition.DiscordWebhook{
			Username: "Applications",
			Embeds: []definition.DiscordWebhookEmbed{
				{
					Title:     "New Application for " + application.Username,
					Timestamp: time.Now().Format(time.RFC3339),
					Color:     0x616161,
					Fields: []definition.DiscordWebhookEmbedField{
						{
							Name:   "ID",
							Value:  id,
							Inline: false,
						},
						{
							Name:   "Username",
							Value:  application.Username,
							Inline: true,
						},
						{
							Name:   "Age",
							Value:  strconv.Itoa(application.Age),
							Inline: true,
						},
						{
							Name:   "Where did you find Orchid?",
							Value:  application.WhereDidYouFindTheServer,
							Inline: false,
						},
						{
							Name:   "Tell us a little about yourself",
							Value:  application.Bio,
							Inline: false,
						},
					},
					Author: &definition.DiscordWebhookEmbedAuthor{
						Name:    application.Username,
						IconUrl: "https://mc-heads.net/head/" + application.Uuid,
					},
				},
			},
		})
	}()
}
