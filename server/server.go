package main

import (
	"apply/route"
	"apply/util"
	"fmt"
	"net/http"
	"strconv"
)

func startServer() {
	http.HandleFunc("/api/apply", route.Apply)
	http.HandleFunc("/api/approved", route.Approved)
	http.HandleFunc("/api/status/{username}", route.CheckStatus)
	http.HandleFunc("/api/invite", route.CreateInvite)
	http.HandleFunc("/api/info", route.Info)

	fmt.Println("Starting HTTP server at 0.0.0.0.0:" + strconv.Itoa(util.Config.Port))
	util.MustNotError(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(util.Config.Port), nil))
}
