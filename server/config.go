package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"orchidmc.org/apply/server/definition"
	"orchidmc.org/apply/server/util"

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
				panic(marshalErr)
			}

			configBytes = marshal

			writeErr := os.WriteFile("config.json", configBytes, 0660)
			if writeErr != nil {
				panic(writeErr)
			}

			return
		}

		panic(openErr)
	}

	bytes, readErr := io.ReadAll(jsonData)
	if readErr != nil {
		panic(readErr)
	}

	defer util.MustNotError(jsonData.Close())

	unmarshalErr := json.Unmarshal(bytes, &util.Config)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
}
