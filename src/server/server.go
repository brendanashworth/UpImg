package server

import (
	"fmt"
	"net/http"
)

// The index page. This also handles 
func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<html><head><title>UpImg</title></head><body><h1>UPIMG UPLOAD UR FILEZ</body></html>")
}

// The upload page.
func UploadHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "<html><head><title>uploadz</title></head><body>dez uploads</body></html>")
}

// Starts the UpImgServer.
func Start(port string) (output string) {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadHandler)

	// Listen and serve.
	http.ListenAndServe(port, nil)

	return "Started UpImg on port "+port
}