package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	speakersSection := findSection(doc)
	fmt.Println(speakersSection.Data)
}

func findSection(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	if n.Type == html.ElementNode && n.Data == "section" {
		return n
	}
	if x := findSection(n.FirstChild); x != nil {
		return x
	}
	return findSection(n.NextSibling)
}
