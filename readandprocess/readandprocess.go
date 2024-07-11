package readandprocess

import (
	"fmt"
	"os"
	"strings"

	"github.com/johneliud/ASCII-Art-Web/checkfilevalidity"
)

// Function ReadAndProcess accepts a banner file string and returns a split slice of string. The function reads and checks for file integrity before returning a split version of it.
func ReadAndProcess(bannerFile string) []string {
	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Printf("Error reading %v file: %v\n", bannerFileData, err)
		os.Exit(1)
	}

	fileHash := checkfilevalidity.CheckFileValidity(bannerFileData)

	switch bannerFile {
	case "standard.txt":
		if fileHash != "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf" {
			fmt.Printf("Invalid hash for %v file. File might be corrupted.\n", bannerFile)
			os.Exit(1)
		}
	case "shadow.txt":
		if fileHash != "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73" {
			fmt.Printf("Invalid hash for %v file. File might be corrupted.\n", bannerFile)
			os.Exit(1)
		}
	case "thinkertoy.txt":
		if fileHash != "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3" {
			fmt.Printf("Invalid hash for %v file. File might be corrupted.\n", bannerFile)
			os.Exit(1)
		}
	}

	// Replace all occurrences of "\r\n" with "\n" to normalize line endings
	splitBannerFileData := strings.ReplaceAll(string(bannerFileData), "\r\n", "\n")
	return strings.Split(splitBannerFileData, "\n")
}
