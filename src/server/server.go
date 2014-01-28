package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

// The index page. This also handles 
func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	// Load the /static/index.html template page.
	content, err := ioutil.ReadFile("static/index.html")
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

// The site.css file.
func StyleHandler(writer http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadFile("static/site.css")
	if err != nil {
		fmt.Println("Error occurred while loading static/site.css: " + err.Error())
		return
	}

	writer.Write(content)
}

// Starts the UpImgServer.
func Start(port string) {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/site.css", StyleHandler)

	// Listen and serve.
	http.ListenAndServe(port, nil)
}