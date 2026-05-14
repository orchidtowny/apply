package main

import (
	"apply/route"
	"apply/util"
	"fmt"
	"net/http"
	"strconv"
)

func startServer() {
	http.HandleFunc("/api/info", route.Info)

	http.HandleFunc("/api/apply", route.Apply)
	http.HandleFunc("/api/status/{username}", route.CheckStatus)

	// Mod Routes
	http.HandleFunc("/api/approved", route.Approved)
	http.HandleFunc("/api/approve/{id}", route.Approve)
	http.HandleFunc("/api/reject/{id}", route.Reject)

	fmt.Println("Starting HTTP server at 0.0.0.0.0:" + strconv.Itoa(util.Config.Port))
	util.MustNotError(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(util.Config.Port), nil))
}
