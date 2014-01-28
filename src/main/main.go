package main

import (
	"fmt"
	"server"
)

func main() {
	fmt.Println("Starting UpImg server on :8080")
	
	// Create the server and start it.
	server.Start(":8080")
}