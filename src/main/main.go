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
	rumtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Starting UpImg server on :8080")

	// Create the database.
	db := backend.ConnectDatabase("localhost", "upimg", "test", "")


	// Create the server and start it.
	*server := server.NewUpImgServer(":8080", &db)
	*server.Start()
}