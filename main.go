package main

import (
	"fmt"
	"os"

	"github.com/johneliud/ASCII-Art-Web/serverhandlers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [PORT]")
		return
	}
	serverhandlers.StartServer()
}
