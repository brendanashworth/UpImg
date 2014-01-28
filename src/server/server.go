package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

// The index page. This also handles 
func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	// Load the /static/index.html template page.
	content, err := ioutil.ReadFile("/static/index.html")
	if err != nil {
		fmt.Println("Error occurred while loading static/index.html: " + err.Error())
		return
	}

	writer.Write(content)
}

// The upload page.
func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<html><head><title>Upload Page</title></head><body>Dat Upload Page</body></html>")
}

// Starts the UpImgServer.
func Start(port string) (output string) {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadHandler)

	// Listen and serve.
	http.ListenAndServe(port, nil)

	return "Started UpImg on port "+port
}