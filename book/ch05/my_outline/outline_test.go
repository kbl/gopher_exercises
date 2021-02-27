package my_outline

import (
	"golang.org/x/net/html"
	"os"
	"strings"
	"testing"
)

func TestOutline(t *testing.T) {
	file, err := os.Open("./index.html")
	if err != nil {
		t.Error("error during opening file")
	}
	parsedHtml := Outline(file)
	_, err = html.Parse(strings.NewReader(parsedHtml))
	if err != nil {
		t.Error("html.Parse(pageHtml) = nil")
	}
}
