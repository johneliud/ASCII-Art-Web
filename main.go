package main

import (
	"fmt"
	"os"

	"github.com/johneliud/ASCII-Art-Web/serverhandlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	serverhandlers.StartServer()
}
