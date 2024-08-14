package serverhandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/johneliud/ASCII-Art-Web/printart"
	"github.com/johneliud/ASCII-Art-Web/readandprocess"
)

type Page struct {
	AsciiArt, Input string
}

// Function HomeHandler serves the home page of the ASCII Art Web application.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		serve404Page(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		serve404Page(w, r)
		return
	}
	tmpl.Execute(w, nil)
}

// Function AsciiArtHandler accepts POST requests at the "/ascii-art" path and generates ASCII art based on the provided input string and banner.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		serve405Page(w, r)
		return
	}

	inputString := r.FormValue("input")
	bannerFile := r.FormValue("banner")

	if inputString == "" {
		serve400Page(w, r)
		return
	}

	if bannerFile == "" {
		bannerFile = "standard"
	}

	bannerFilePath := fmt.Sprintf("banners/%s.txt", bannerFile)
	bannerFileSlice, err := readandprocess.ReadAndProcess(bannerFilePath)
	if err != nil {
		serve500Page(w, r)
		return
	}

	var output strings.Builder
	err = printart.PrintArt(bannerFileSlice, inputString, &output)
	if err != nil {
		serve400Page(w, r)
		return
	}

	data := Page{
		Input:    inputString,
		AsciiArt: output.String(),
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		serve404Page(w, r)
		return
	}
	tmpl.Execute(w, data)
}

// Function ExportArtToFile listens for Get requests on the download button and writes the generated art to a file.
func ExportArtToFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		serve405Page(w, r)
		return
	}

	// Extract art from query
	generatedArt := r.URL.Query().Get("ascii")
	if generatedArt == "" {
		serve400Page(w, r)
		return
	}

	generatedArtInBytes := []byte(generatedArt)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(generatedArtInBytes)))

	_, err := w.Write(generatedArtInBytes)
	if err != nil {
		log.Printf("Error writing ASCII art to response: %v\n", err)
		serve500Page(w, r)
		return
	}
}
