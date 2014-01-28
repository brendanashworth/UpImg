package main

import (
	"fmt"
	"server"
)

func main() {
	
	// Create the server and start it.
	output := server.Start(":8080")

	fmt.Println(output)
}