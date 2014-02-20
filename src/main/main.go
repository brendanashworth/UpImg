package main

import (
	"fmt"
	"runtime"
	// These come from our packages.
	"server"
)

func main() {
	// Setup Go stuff
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Starting UpImg server on :8080...")

	// Create the server and start it.
	server := server.NewServer(":8080")
	server.Start()
}