package serverhandlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/johneliud/ASCII-Art-Web/printart"
	"github.com/johneliud/ASCII-Art-Web/readandprocess"
)

type Page struct {
	AsciiArt string
}

/*
Function HomeHandler is an HTTP handler function that serves the home page of the ASCII Art Web application. If the root path is requested, it attempts to parse and execute the "static/index.html" template file.
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Checks if the requested URL path is the root ("/") and returns a 404 Not Found status if not
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Render the HTML template file and write the output to the HTTP response writer
	tmpl.Execute(w, nil)
}

/*
Function AsciiArtHandler is a HTTP handler function that processes ASCII art requests. It accepts POST requests at the "/ascii-art" path and generates ASCII art based on the provided input string and banner.
*/
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Checks if the request method is POST. If not, it returns a 405 Method Not Allowed status
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieves the input string and banner file from the request form data
	inputString := r.FormValue("input")
	bannerFile := r.FormValue("banner")

	// Validates the input string. If it is empty, it returns a 400 Bad Request status.
	if inputString == "" {
		http.Error(w, "Input string is required", http.StatusBadRequest)
		return
	}

	// Sets the default banner file if none is provided
	if bannerFile == "" {
		bannerFile = "standard"
	}

	bannerFilePath := fmt.Sprintf("banners/%s.txt", bannerFile)
	bannerFileSlice := readandprocess.ReadAndProcess(bannerFilePath)

	var output strings.Builder
	// Generates ASCII art using the PrintArt function and writes the output to a strings.Builder
	printart.PrintArt(bannerFileSlice, inputString, &output)

	// Creates a Page struct with the generated ASCII art
	data := Page{
		AsciiArt: output.String(),
	}

	// Parses and executes the "static/index.html" template file with the Page data and returns a 500 Internal Server Error status if the template cannot be found
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

/*
Function StartServer sets up and starts the HTTP server for the ASCII Art Web application. It serves static files from the "static" directory, and registers handlers for the root ("/") and "/ascii-art" paths. The server listens on port 8080 and prints a message indicating its running status.
*/
func StartServer() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
