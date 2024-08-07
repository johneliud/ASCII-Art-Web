package readandprocess

import (
	"fmt"
	"os"
	"strings"

	"github.com/johneliud/ASCII-Art-Web/checkfilevalidity"
	"github.com/johneliud/ASCII-Art-Web/downloadbanner"
)

// Function ReadAndProcess accepts a banner file string and returns a split slice of string. The function reads and checks for file integrity, and downloads a new file if needed before returning a split version of it.
func ReadAndProcess(bannerFile string) ([]string, error) {
	bannerFileData, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, fmt.Errorf("error reading %v file: %v", bannerFileData, err)
	}

	fileHash := checkfilevalidity.CheckFileValidity(bannerFileData)

	var expectedHashValue string
	switch bannerFile {
	case "banners/standard.txt":
		expectedHashValue = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	case "banners/shadow.txt":
		expectedHashValue = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	case "banners/thinkertoy.txt":
		expectedHashValue = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	default:
		return nil, fmt.Errorf("%v is an invalid banner file", bannerFile)
	}

	if fileHash != expectedHashValue {
		fmt.Printf("File %v might be corrupted. Deleting and downloading a new file.\n", bannerFile)
		err = os.Remove(bannerFile)
		if err != nil {
			return nil, fmt.Errorf("an unexpected error occurred while deleting the corrupted file. %v", err)
		}

		err = downloadbanner.DownloadBannerFile(bannerFile)
		if err != nil {
			return nil, fmt.Errorf("an unexpected error occurred while downloading the new file. %v", err)
		}

		bannerFileData, err = os.ReadFile(bannerFile)
		if err != nil {
			return nil, fmt.Errorf("error reading %v file after downloading: %v", bannerFile, err)
		}
	}

	splitBannerFileData := strings.ReplaceAll(string(bannerFileData), "\r\n", "\n")
	return strings.Split(splitBannerFileData, "\n"), nil
}
