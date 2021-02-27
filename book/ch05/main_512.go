package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/ex512"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	// ex512.ForEachNode(doc, ex512.StartElement, ex512.EndElement)
	ex512.Outline(doc)
}
