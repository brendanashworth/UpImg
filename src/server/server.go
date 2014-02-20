package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"math/rand"
	"strconv"
)

type UpImgServer struct {
	port string
}

func NewServer(port string) *UpImgServer {
	return &UpImgServer{
		port: port,
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
			http.Redirect(writer, request, `/`, 302)
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
	// make sure it was POST
	if request.Method != "POST" {
		return
	}

	file, _, err := request.FormFile("file")
	if err != nil {
		fmt.Println("Error uploading file #1: " + err.Error())
		return
	}

	// now that hes uploaded it, get the data
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error uploading file #2: " + err.Error())
	}

	name := GenerateNewURL()

	// now try and save the file.
	err = ioutil.WriteFile("images/" + name + ".png", data, 0777)
	if err != nil {
		fmt.Println("Error uploading file #3: " + err.Error())
		http.Redirect(writer, request, `/`, 302) // redirect him to home
	}

	http.Redirect(writer, request, "/" + name, 302)
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


// utilities

func GenerateNewURL() string {
	var chars = ""
	for i := 0; i < 7; i = i + 1 {
		chars = chars + strconv.Itoa(rand.Intn(10))
	}

	return chars
}

