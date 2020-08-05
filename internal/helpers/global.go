package helpers

import (
	"net/url"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

// IsURL function - validate a URL.
func IsURL(value string) bool {
	result, err := url.Parse(value)
	return err == nil && result.Scheme != "" && result.Host != ""
}

// FileExists function - check with a file exist in system.
func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Empty function - check if a string is empty.
func Empty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// ASCIIMessage function - print a message using ASCII Art.
func ASCIIMessage(content, font, color string) {
	var c string
	if !Empty(color) {
		c = color
	} else {
		c = "green"
	}
	if !Empty(font) {
		figure.NewColorFigure(content, font, c, true).Print()
	} else {
		figure.NewColorFigure(content, "", c, true).Print()
	}
}
