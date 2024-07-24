package downloadbanner

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


// Function DownloadBannerFile accepts a file path and downloads the corresponding ASCII art banner file from GitHub.
func DownloadBannerFile(filePath string) error {
	var downloadURL string

	switch filePath {
	case "banners/standard.txt":
		downloadURL = "https://github.com/johneliud/ASCII-Art-Web/raw/main/banners/standard.txt"
	case "banners/shadow.txt":
		downloadURL = "https://github.com/johneliud/ASCII-Art-Web/raw/main/banners/shadow.txt"
	case "banners/thinkertoy.txt":
		downloadURL = "https://github.com/johneliud/ASCII-Art-Web/raw/main/banners/thinkertoy.txt"
	default:
		return fmt.Errorf("invalid file path: %v", filePath)
	}

	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("error downloading file from %v: %v", downloadURL, err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file %v: %v", filePath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing to file %v: %v", filePath, err)
	}
	return nil
}
