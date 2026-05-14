package main

import (
	"apply/definition"
	"apply/util"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

func loadConfig() {
	jsonData, openErr := os.Open("config.json")
	if openErr != nil {
		if os.IsNotExist(openErr) {
			fmt.Println("Config file does not exist, writing default")

			var configBytes []byte

			marshal, marshalErr := json.Marshal(definition.Config{
				Port:              8070,
				ServerIp:          "play.orchidmc.org",
				ApiKey:            uuid.New().String(),
				DiscordWebhookUrl: "",
				Rules:             map[string]string{},
			})
			if marshalErr != nil {
				fmt.Println("Error encoding default config:", marshalErr)
				return
			}

			configBytes = marshal

			writeErr := os.WriteFile("config.json", configBytes, 0660)
			if writeErr != nil {
				fmt.Println("Error creating config:", writeErr)
				return
			}

			return
		}

		fmt.Println("Error loading config:", openErr)
		return
	}

	bytes, readErr := io.ReadAll(jsonData)
	if readErr != nil {
		fmt.Println("Error reading config:", readErr)
		return
	}

	defer util.MustNotError(jsonData.Close())

	unmarshalErr := json.Unmarshal(bytes, &config)
	if unmarshalErr != nil {
		fmt.Println("Error parsing config:", unmarshalErr)
		return
	}
}
