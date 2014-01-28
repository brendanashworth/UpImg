package main

import (
	"fmt"
	"github.com/boboman13/UpImg/server"
)

func main() {
	
	// Create the server and start it.
	server := UpImgServer{}
	output := server.Start(":80")

	fmt.Println(output)
}