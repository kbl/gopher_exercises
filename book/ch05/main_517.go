package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/ex517"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementsbytagname: %v\n", err)
		os.Exit(1)
	}
	for _, e := range ex517.ElementsByTagName(doc, "h1", "h2", "h3", "h4") {
		fmt.Println(e.Data)
	}
}
