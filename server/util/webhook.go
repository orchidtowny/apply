package util

import (
	"applyServer/definition"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendWebhookMessage(msg definition.DiscordWebhook) {
	body, webhookBodyErr := json.Marshal(msg)
	if webhookBodyErr != nil {
		fmt.Printf("Failed to marshal webhook body: " + webhookBodyErr.Error())
		return
	}

	post, postErr := http.Post(
		Config.DiscordWebhookUrl,
		"application/json",
		bytes.NewBuffer(body),
	)
	if postErr != nil {
		fmt.Printf("Failed to post webhook: " + postErr.Error())
		return
	}

	if post.StatusCode != http.StatusOK {
		response, readWebhookErr := io.ReadAll(post.Body)
		if readWebhookErr != nil {
			fmt.Printf("Failed to read webhook response: " + readWebhookErr.Error())
			return
		}

		text := string(response)
		fmt.Printf("Failed to post webhook status: " + post.Status + ": " + text)

		return
	}
}
