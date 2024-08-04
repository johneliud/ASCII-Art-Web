package serverhandlers

import (
	"html/template"
	"log"
	"net/http"
)

func serveHttpError(w http.ResponseWriter, r *http.Request, statusCode int, templatePath string, statusMessage string) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, statusMessage, statusCode)
		log.Printf("Error parsing %v: %v\n", templatePath, err)
		return
	}
	w.WriteHeader(statusCode)
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, statusMessage, statusCode)
		log.Printf("Error executing %v: %v\n", templatePath, err)
		return
	}
}

func serve400Page(w http.ResponseWriter, r *http.Request) {
	serveHttpError(w, r, http.StatusBadRequest, "templates/page400.html", "Bad Request")
}

func serve404Page(w http.ResponseWriter, r *http.Request) {
	serveHttpError(w, r, http.StatusBadRequest, "templates/page404.html", "Not Found")
}

func serve405Page(w http.ResponseWriter, r *http.Request) {
	serveHttpError(w, r, http.StatusBadRequest, "templates/page405.html", "Method Not Allowed")
}

func serve500Page(w http.ResponseWriter, r *http.Request) {
	serveHttpError(w, r, http.StatusBadRequest, "templates/page500.html", "An Unexpected Error Occurred. Try Again Later")
}
