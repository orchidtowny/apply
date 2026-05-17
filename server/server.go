package main

import (
	"fmt"
	"net/http"
	"strconv"

	"orchidmc.org/apply/server/route"
	"orchidmc.org/apply/server/util"
)

func startServer() {
	http.HandleFunc("/api/info", route.Info)

	http.HandleFunc("/api/apply", route.Apply)
	http.HandleFunc("/api/status/{username}", route.CheckStatus)

	// Mod Routes)
	http.HandleFunc("/api/approved", route.Approved)
	http.HandleFunc("/api/add", route.ManuallyAdd)
	http.HandleFunc("/api/approve/{id}", route.Approve)
	http.HandleFunc("/api/reject/{id}", route.Reject)

	// Frontend
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)
	http.Handle("/status", fs)
	http.Handle("/admin", fs)
	http.Handle("/assets/*", fs)

	fmt.Println("Starting HTTP server at 0.0.0.0.0:" + strconv.Itoa(util.Config.Port))
	util.MustNotError(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(util.Config.Port), nil))
}
