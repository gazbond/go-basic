package main

import (
	"github.com/gazbond/go-basic/server"
)

// Run server
func main() {

	server.LoadConfig()
	server.Start()
}
