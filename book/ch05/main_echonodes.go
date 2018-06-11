package main

import (
	"book/ch05/echonodes"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	echonodes.EchoNodes(doc, "")
}
