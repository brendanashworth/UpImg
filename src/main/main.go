package main

import (
	"fmt"
	"runtime"
	// These come from our packages.
	"server"
	"backend"
)

func main() {
	// Setup Go stuff
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Starting UpImg server on :8080")

	// Create the database.
	db := backend.ConnectDatabase("localhost", "upimg", "test", "")

	// Create the server and start it.
	server := server.NewServer(":8080", db)
	server.Start()
}