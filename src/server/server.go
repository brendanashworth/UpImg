package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"backend"
	"database/sql"
)

type UpImgServer struct {
	port string
	db *sql.DB
}

func NewServer(port string, db *sql.DB) (server *UpImgServer) {
	*server = &UpImgServer{
		port: port,
		db: *db,
	}

	return *server
}

// Starts the UpImgServer.
func (server *UpImgServer) Start(port string) {
	http.HandleFunc("/", *server.IndexHandler)
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/site.css", StyleHandler)

	// Listen and serve.
	http.ListenAndServe(port, nil)
}

// The index page. This also handles 
func (server *UpImgServer) IndexHandler(writer http.ResponseWriter, request *http.Request) {
	if(len(*request.URL.Path) > 0) {
		//img := backend.GetImage(*request.url.Path)
		//writer.Write(img)

		return
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