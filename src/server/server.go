package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"backend"
	"regexp"
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
	if len(request.URL.Path) == 8 {
		// make sure the url is alphanumeric
		reg, err := regexp.Compile("[^A-Za-z0-9]+")
		if err != nil {
			fmt.Println("Error occurred while compiling regex: " + err.Error())
			return
		}

		url := reg.ReplaceAllString(request.URL.Path[1:8], "")
		img := "images/" + url + ".png"

		// double check the length is still 8
		if len(url) != 7 {
			return
		}

		data, err := ioutil.ReadFile(img)
		if err != nil {
			fmt.Println("Error occurred while reading file: [" + img + "] " + err.Error())
			return
		}

		writer.Write(data)
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
	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("Error uploading file: " + err.Error())
		return
	}

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