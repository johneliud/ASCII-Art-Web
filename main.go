package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/johneliud/ASCII-Art-Web/serverhandlers"
)

func main() {
	var (
		port          int
		listeningPort string
	)

	switch len(os.Args) {
	case 1:
		port = 8080
		listeningPort = ":" + strconv.Itoa(port)
	case 2:
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
		listeningPort = ":" + strconv.Itoa(port)
	default:
		fmt.Println("Incorrect number of arguments passed.")
		return
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.HandleFunc("/", serverhandlers.HomeHandler)
	http.HandleFunc("/ascii-art", serverhandlers.AsciiArtHandler)

	log.Printf("Server running on http://localhost%v\n", listeningPort)
	http.ListenAndServe(listeningPort, nil)
}
