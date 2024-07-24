package serverhandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/johneliud/ASCII-Art-Web/printart"
	"github.com/johneliud/ASCII-Art-Web/readandprocess"
)

type Page struct {
	AsciiArt, Input string
}

// Function HomeHandler is an HTTP handler function that serves the home page of the ASCII Art Web application. If the root path is requested, it attempts to parse and execute the "templates/index.html" template file.
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

// Function AsciiArtHandler is a HTTP handler function that processes ASCII art requests. It accepts POST requests at the "/ascii-art" path and generates ASCII art based on the provided input string and banner.
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
