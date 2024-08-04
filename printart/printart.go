package printart

import (
	"fmt"
	"strings"
)

// PrintArt takes a slice of strings representing ASCII art, an input string, and a strings.Builder. It prints the input string as ASCII art using the provided ASCII art banner. The ASCII art is constructed by replacing certain characters in the input string with their corresponding ASCII representations.
func PrintArt(bannerFileSlice []string, inputString string, output *strings.Builder) error {
	switch inputString {
	case "":
		return nil
	}

	// Handle unprintable sequences
	unprintableSequences := []string{"\\a", "\\b", "\\v", "\\f", "\\r"}
	for _, seq := range unprintableSequences {
		if strings.Contains(inputString, seq) {
			return fmt.Errorf("input string contains an unprintable sequence")
		}
	}
	splitArguments := strings.Split(inputString, "\r\n")

	// Handle characters not present in the ASCII manual
	for _, argument := range splitArguments {
		for _, char := range argument {
			if char < 32 || char > 126 {
				return fmt.Errorf("input string contains a character absent in the ASCII manual")
			}
		}
	}

	for _, text := range splitArguments {
		if text == "" {
			output.WriteString("\n")
			continue
		}

		asciiHeight := 8
		for i := 0; i < asciiHeight; i++ {
			for _, char := range text {
				startingIndex := int(char-32)*9 + 1
				output.WriteString(bannerFileSlice[startingIndex+i])
			}
			output.WriteString("\n")
		}
	}
	return nil
}
