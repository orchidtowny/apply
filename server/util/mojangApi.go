package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"orchidmc.org/apply/server/definition"
)

func GetPlayerUuid(username string) string {
	get, getErr := http.Get("https://api.mojang.com/users/profiles/minecraft/" + username)
	if getErr != nil {
		fmt.Println("Error getting player", getErr.Error())
		return ""
	}

	bytes, readErr := io.ReadAll(get.Body)
	if readErr != nil {
		fmt.Println("Error reading player", readErr.Error())
		return ""
	}

	var lookup definition.MojangPlayerLookup
	unmarshalErr := json.Unmarshal(bytes, &lookup)
	if unmarshalErr != nil {
		fmt.Println("Error unmarshaling player", unmarshalErr.Error())
		return ""
	}

	return lookup.Id
}
