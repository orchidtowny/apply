package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// in minutes
const sync_interval = 15

var api_url string = "https://apply.orchidmc.org"
var api_key string
var whitelist_file_location string

var client = &http.Client{}

func main() {
	fmt.Println("Starting sync tool...")

	api_url = os.Getenv("SYNC_API_URL")
	api_key = os.Getenv("SYNC_API_KEY")
	whitelist_file_location = os.Getenv("SYNC_WHITELIST_FILE_LOCATION")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(sync_interval * time.Minute)
	done := make(chan bool)

	go func() {
		for {
			sync()

			select {
			case <-ticker.C:
				continue
			case <-done:
				return
			}
		}
	}()

	<-ctx.Done()

	ticker.Stop()
	done <- true
}

func sync() {
	fmt.Println("Syncing...")

	// Getting data

	req, reqErr := http.NewRequest("GET", api_url+"/api/approved", nil)
	if reqErr != nil {
		fmt.Println("Errored on creating request: " + reqErr.Error())
		return
	}

	req.Header.Add("Authorization", "Bearer "+api_key)

	response, getErr := client.Do(req)
	if getErr != nil {
		fmt.Println("Errored on request: " + getErr.Error())
		return
	}

	bytes, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println("Error reading response: " + readErr.Error())
		return
	}

	var data []MinecraftPlayer
	unmarshalErr := json.Unmarshal(bytes, &data)
	if unmarshalErr != nil {
		fmt.Println("Failed parsing response: " + unmarshalErr.Error())
		return
	}

	fmt.Println("Got " + strconv.Itoa(len(data)) + " approved players, writing to " + whitelist_file_location)

	// Writing

	writableData, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		fmt.Println("Failed formatting writable data: " + marshalErr.Error())
		return
	}

	writeErr := os.WriteFile(whitelist_file_location, writableData, 0660)
	if writeErr != nil {
		fmt.Println("Failed writing data: " + writeErr.Error())
		return
	}

	fmt.Println("Finished sync!")
}
