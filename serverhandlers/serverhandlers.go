package serverhandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/johneliud/ASCII-Art-Web/printart"
	"github.com/johneliud/ASCII-Art-Web/readandprocess"
)

type Page struct {
	AsciiArt, Input string
}

/*
Function HomeHandler is an HTTP handler function that serves the home page of the ASCII Art Web application. If the root path is requested, it attempts to parse and execute the "templates/index.html" template file.
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "an unexpected error occurred. please try again later", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

/*
Function AsciiArtHandler is a HTTP handler function that processes ASCII art requests. It accepts POST requests at the "/ascii-art" path and generates ASCII art based on the provided input string and banner.
*/
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, " 405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieves the input string and banner file from the request form data
	inputString := r.FormValue("input")
	bannerFile := r.FormValue("banner")
	
	if inputString == "" {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}
	
	// Sets the default banner selection to be standard if not specified/selected
	if bannerFile == "" {
		bannerFile = "standard"
	}

	bannerFilePath := fmt.Sprintf("banners/%s.txt", bannerFile)
	bannerFileSlice, err := readandprocess.ReadAndProcess(bannerFilePath)
	if err != nil {
		http.Error(w, "an unexpected error occurred. please try again later", http.StatusInternalServerError)
		return
	}

	var output strings.Builder
	err = printart.PrintArt(bannerFileSlice, inputString, &output)
	if err != nil {
		http.Error(w, "an unexpected error occurred. please try again later", http.StatusInternalServerError)
		return
	}

	data := Page{
		Input:    inputString,
		AsciiArt: output.String(),
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "an unexpected error occurred. please try again later", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

/*
Function StartServer sets up and starts the HTTP server for the ASCII Art Web application. It serves static files from the "static" directory, and registers handlers for the root ("/") and "/ascii-art" paths.
*/
func StartServer() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)
	
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Cannot convert %v to integer: %v\n", os.Args[1], err)
		return
	}

	// Check if portNum is within a valid range
	if port < 1024 || port > 65535 {
		fmt.Println("Port to running the server should be within the range 1024 to 65535")
		return
	}
	listeningPort := ":" + strconv.Itoa(port)
	log.Printf("Server running on http://localhost%v\n", listeningPort)
	http.ListenAndServe(listeningPort, nil)
}
