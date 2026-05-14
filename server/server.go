package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	serveErr := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(config.Port), nil)
	if serveErr != nil {
		fmt.Println("Error starting server:", serveErr)
		return
	}
}
