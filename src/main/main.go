package main

import (
	"fmt"
	"server"
	"runtime"
)

func main() {
	// Setup Go stuff
	rumtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Starting UpImg server on :8080")
	
	// Create the server and start it.
	server.Start(":8080")
}