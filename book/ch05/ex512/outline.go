package ex512

import (
	"book/ch05/outline2"
	"fmt"
	"golang.org/x/net/html"
)

func Outline(n *html.Node) {
	var depth int
	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	outline2.ForEachNode(n, startElement, endElement)
}
