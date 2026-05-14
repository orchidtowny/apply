package main

import (
	"apply/definition"
)

var config definition.Config

func main() {
	loadConfig()
	startServer()
}
