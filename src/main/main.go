package main

import (
	"fmt"
	"server"
)

func main() {
	
	// Create the server and start it.
	server := UpImgServer{}
	output := server.Start(":80")

	fmt.Println(output)
}