package main

import (
	"fmt"
	"server"
)

func main() {
	
	// Create the server and start it.
	server := UpImgServer{}
	output := server.Start(":8080")

	fmt.Println(output)
}