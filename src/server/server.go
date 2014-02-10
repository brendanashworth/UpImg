package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"backend"
)

type UpImgServer struct {
	port string
	db *backend.Backend
}

func NewServer(port string, db *backend.Backend) *UpImgServer {
	return &UpImgServer{
		port: port,
		db: db,
	}
}

// Starts the UpImgServer.
func (this *UpImgServer) Start() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/site.css", StyleHandler)

	// Listen and serve.
	http.ListenAndServe(this.port, nil)
}

// The index page. This also handles image requests.
func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	// If the user asked for an image.
	if len(request.URL.Path) == 9 {
		fmt.Fprintf(writer, request.URL.Path[1:8])
		return
	} else if len(request.URL.Path) == 1 {
		// It was `/`, don't redirect.
	} else {
		http.Redirect(writer, request, `/`, 302)
	}

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
	fmt.Fprintf(writer, "Upload complete.")
}

// The site.css file.
func StyleHandler(writer http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadFile("static/site.css")
	if err != nil {
		fmt.Println("Error occurred while loading static/site.css: " + err.Error())
		return
	}

	writer.Header().Set("Content-Type", "text/css")
	writer.Write(content)
}